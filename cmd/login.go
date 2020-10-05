package cmd

import (
	"bufio"
	"fmt"
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
		Token string `json:"token" yaml:"token"`
	}

	// Credentials holds the registry token when a user performs a login.
	Credentials struct {
		Registry Registry `json:"registry" yaml:"registry"`
	}
)

func loginCommand() *cli.Command {
	return &cli.Command{
		Name: "login",
		Usage: `Save an API token from the Atlas registry locally. It a token is not specified, it will
read from stdin.`,
		ArgsUsage: "[token]",
		Action: func(ctx *cli.Context) error {
			var token string

			if ctx.NArg() > 0 {
				token = ctx.Args().Get(0)
			} else {
				reader := bufio.NewReader(os.Stdin)

				// TODO: Add prompt.
				input, err := reader.ReadString('\n')
				if err != nil {
					return fmt.Errorf("failed to read input: %w", err)
				}

				token = strings.TrimSuffix(input, "\n")
			}

			dir := path.Join(os.Getenv("HOME"), ".atlas")
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

			_, _ = color.New(color.FgCyan).Fprintln(os.Stderr, "login token successfully saved!")
			return nil
		},
	}
}
