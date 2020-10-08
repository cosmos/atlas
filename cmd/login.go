package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

type (
	// Registry defines the registries token value when a user performs a login.
	Registry struct {
		Token string `json:"token" toml:"token"`
	}

	// Credentials holds the registry token when a user performs a login.
	Credentials struct {
		Registry Registry `json:"registry" toml:"registry"`
	}
)

// LoginCommand returns a login CLI command handler.
func LoginCommand() *cli.Command {
	return &cli.Command{
		Name: "login",
		Usage: `Save an API token from the Atlas registry locally. It a token is not specified, it will
read from stdin.`,
		ArgsUsage: "[token]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Value:   os.Getenv("HOME"),
				Usage:   "The root directory for Atlas configuration",
			},
		},
		Action: func(ctx *cli.Context) error {
			var token string

			if ctx.NArg() > 0 {
				token = ctx.Args().Get(0)
			} else {
				var reader io.Reader
				if v := ctx.Context.Value(cmdReaderKey); v != nil {
					reader = v.(io.Reader)
				} else {
					reader = os.Stdin
				}

				buffReader := bufio.NewReader(reader)

				// TODO: Add prompt.
				input, err := buffReader.ReadString('\n')
				if err != nil {
					return fmt.Errorf("failed to read input: %w", err)
				}

				token = strings.TrimSuffix(input, "\n")
			}

			dir := path.Join(ctx.String("dir"), ".atlas")
			if err := os.MkdirAll(dir, 0700); err != nil {
				return err
			}

			credsPath := path.Join(dir, "credentials")
			file, err := os.Create(credsPath)
			if err != nil {
				return err
			}

			defer func() {
				_ = file.Close()
			}()

			if err := toml.NewEncoder(file).Encode(Credentials{Registry{Token: token}}); err != nil {
				return err
			}

			_, _ = color.New(color.FgGreen).Fprintf(ctx.App.Writer, "login token successfully saved to %s\n", dir)
			return nil
		},
	}
}

func parseCredentials(credsPath string) (Credentials, error) {
	var creds Credentials

	if _, err := toml.DecodeFile(credsPath, &creds); err != nil {
		return Credentials{}, fmt.Errorf("failed to read credentials: %w", err)
	}

	return creds, nil
}
