package cmd

import (
	"fmt"
	"io"
	"log/syslog"
	"os"
	"strings"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/server"
	"github.com/cosmos/atlas/server/crawl"
)

// StartServerCommand returns a CLI command handler responsible for starting
// the Atlas service daemon.
func StartServerCommand() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Start the atlas server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    config.ConfigPath,
				Aliases: []string{"c"},
				Usage:   "Server configuration file.",
			},
			&cli.StringFlag{
				Name:  config.LogLevel,
				Value: zerolog.InfoLevel.String(),
				Usage: "The server logging level (panic|fatal|error|warn|info|debug|trace)",
			},
			&cli.StringFlag{
				Name:  config.LogFormat,
				Value: "json",
				Usage: "The server logging format (text|json)",
			},
			&cli.StringFlag{
				Name:  config.ListenAddr,
				Value: "localhost:8080",
				Usage: "The server listen address",
			},
			&cli.BoolFlag{
				Name:  config.Dev,
				Value: false,
				Usage: "Enable development settings used for non-production environments",
			},
			&cli.DurationFlag{
				Name:  config.HTTPReadTimeout,
				Value: 15 * time.Second,
				Usage: "Define the HTTP read timeout",
			},
			&cli.DurationFlag{
				Name:  config.HTTPWriteTimeout,
				Value: 15 * time.Second,
				Usage: "Define the HTTP write timeout",
			},
		},
		Action: func(ctx *cli.Context) error {
			konfig, err := ParseServerConfig(ctx)
			if err != nil {
				return err
			}

			logLvl, err := zerolog.ParseLevel(konfig.String(config.LogLevel))
			if err != nil {
				return fmt.Errorf("failed to parse log level: %w", err)
			}

			var logWriter io.Writer
			if strings.ToLower(konfig.String(config.LogFormat)) == "text" {
				logWriter = zerolog.ConsoleWriter{Out: os.Stderr}
			} else {
				logWriter = os.Stderr
			}

			// When not in dev mode, in addition to writing logs to logWriter, also
			// stream logs to the configured syslog sink via TCP using a multi-writer.
			syslogAddr := konfig.String(config.SyslogAddr)
			if !konfig.Bool(config.Dev) && syslogAddr != "" {
				ptWriter, err := syslog.Dial("tcp", syslogAddr, syslog.LOG_EMERG|syslog.LOG_KERN, "atlas")
				if err != nil {
					return err
				}

				logWriter = zerolog.MultiLevelWriter(logWriter, ptWriter)
			}

			logger := zerolog.New(logWriter).Level(logLvl).With().Timestamp().Logger()

			svr, err := server.NewService(logger, konfig)
			if err != nil {
				return err
			}

			// start the service in a separate goroutine
			go func() {
				if err := svr.Start(); err != nil {
					logger.Fatal().Err(err).Msg("failed to start atlas service")
				}
			}()

			crawler, err := crawl.NewCrawler(logger, konfig, svr.GetDB())
			if err != nil {
				return err
			}

			if crawler != nil {
				// start the node crawler in a separate goroutine
				go crawler.Start()
			}

			// trap signals and perform any cleanup
			trapSignal(func() {
				logger.Info().Msg("shuting down...")

				if crawler != nil {
					crawler.Stop()
				}

				svr.Cleanup()
				os.Exit(0)
			})

			// run forever...
			select {}
		},
	}
}

// ParseServerConfig returns a server configuration, given a command Context,
// by parsing the following in order of precedence:
//
// - flags
// - environment variables
// - configuration file (TOML)
func ParseServerConfig(ctx *cli.Context) (*koanf.Koanf, error) {
	konfig := koanf.New(".")

	// load from file first (if provided)
	if configPath := ctx.String(config.ConfigPath); len(configPath) != 0 {
		if err := konfig.Load(file.Provider(configPath), toml.Parser()); err != nil {
			return nil, err
		}
	}

	// load from environment variables
	if err := konfig.Load(env.Provider("ATLAS_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(strings.TrimPrefix(s, "ATLAS_")), "_", ".", -1)
	}), nil); err != nil {
		return nil, err
	}

	// finally, load from command line flags
	if err := konfig.Load(NewCLIFlagProvider(ctx, ".", konfig), nil); err != nil {
		return nil, err
	}

	return konfig, nil
}
