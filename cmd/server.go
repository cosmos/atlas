package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/rs/zerolog"
	"github.com/urfave/cli/v2"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/server"
)

// StartServerCommand returns a CLI command handler responsible for starting
// the Atlas service daemon.
func StartServerCommand() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Start the atlas server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    config.FlagConfig,
				Aliases: []string{"c"},
				Usage:   "Server configuration file.",
			},
			&cli.StringFlag{
				Name:  config.FlagLogLevel,
				Value: zerolog.InfoLevel.String(),
				Usage: "The server logging level (panic|fatal|error|warn|info|debug|trace)",
			},
			&cli.StringFlag{
				Name:  config.FlagLogFormat,
				Value: "json",
				Usage: "The server logging format (text|json)",
			},
			&cli.StringFlag{
				Name:  config.FlagListenAddr,
				Value: "localhost:8080",
				Usage: "The server listen address",
			},
			&cli.BoolFlag{
				Name:  config.FlagDev,
				Value: false,
				Usage: "Enable development settings used for non-production environments",
			},
			&cli.DurationFlag{
				Name:  config.FlagHTTPReadTimeout,
				Value: 15 * time.Second,
				Usage: "Define the HTTP read timeout",
			},
			&cli.DurationFlag{
				Name:  config.FlagHTTPWriteTimeout,
				Value: 15 * time.Second,
				Usage: "Define the HTTP write timeout",
			},
		},
		Action: func(ctx *cli.Context) error {
			// Read configuration in order of precedence:
			//
			// - flags
			// - environment variables
			// - configuration file
			konfig := koanf.New(".")

			// load from file first (if provided)
			if configPath := ctx.String(config.FlagConfig); len(configPath) != 0 {
				if err := konfig.Load(file.Provider(configPath), yaml.Parser()); err != nil {
					return err
				}
			}

			// load from environment variables
			if err := konfig.Load(env.Provider("ATLAS_", ".", func(s string) string {
				return strings.Replace(strings.ToLower(strings.TrimPrefix(s, "ATLAS_")), "_", ".", -1)
			}), nil); err != nil {
				return err
			}

			// finally, load from command line flags
			if err := konfig.Load(NewCLIFlagProvider(ctx, ".", konfig), nil); err != nil {
				return err
			}

			logLvl, err := zerolog.ParseLevel(konfig.String(config.FlagLogLevel))
			if err != nil {
				return fmt.Errorf("failed to parse log level: %w", err)
			}

			var logWriter io.Writer
			if strings.ToLower(konfig.String(config.FlagLogFormat)) == "text" {
				logWriter = zerolog.ConsoleWriter{Out: os.Stderr}
			} else {
				logWriter = os.Stderr
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

			// trap signals and perform any cleanup
			trapSignal(func() {
				logger.Info().Msg("shuting down...")
				svr.Cleanup()
				os.Exit(0)
			})

			// run forever...
			select {}
		},
	}
}
