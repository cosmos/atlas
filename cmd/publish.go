package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/urfave/cli/v2"

	"github.com/cosmos/atlas/server"
)

var (
	validate = validator.New()

	client = http.Client{
		Timeout: 15 * time.Second,
	}
)

// PublishCommand returns a CLI command handler responsible for publishing
// Cosmos SDK modules to the Atlas registry.
func PublishCommand() *cli.Command {
	return &cli.Command{
		Name:  "publish",
		Usage: `Publish a Cosmos SDK module to the Atlas registry.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Value:   os.Getenv("HOME"),
				Usage:   "The root directory for Atlas configuration",
			},
			&cli.StringFlag{
				Name:    "manifest",
				Aliases: []string{"m"},
				Value:   path.Join(mustGetwd(), "atlas.toml"),
				Usage:   "The path to the Cosmos SDK module manifest",
			},
			&cli.StringFlag{
				Name:    "registry",
				Aliases: []string{"r"},
				Value:   "https://atlas.cosmos.network",
				Usage:   "The Atlas registry address",
			},
			&cli.BoolFlag{
				Name:  "verify-only",
				Value: false,
				Usage: "Only verify the module manifest without publishing",
			},
		},
		Action: func(ctx *cli.Context) error {
			var manifest server.Manifest

			// fetch and decode the manifest
			manifestPath := ctx.String("manifest")
			if _, err := toml.DecodeFile(manifestPath, &manifest); err != nil {
				return fmt.Errorf("failed to read manifest: %w", err)
			}

			// verify the contents if requested
			if ctx.Bool("verify-only") {
				if err := validate.Struct(manifest); err != nil {
					return fmt.Errorf("failed to verify manifest: %w", server.TransformValidationError(err))
				}

				_, _ = color.New(color.FgGreen).Fprintln(ctx.App.Writer, "manifest successfully verified!")
				return nil
			}

			// fetch the user token from configuration
			dir := path.Join(ctx.String("dir"), ".atlas")
			credsPath := path.Join(dir, "credentials")

			credentials, err := parseCredentials(credsPath)
			if err != nil {
				return err
			}

			// make the API request
			bodyBz, err := json.Marshal(manifest)
			if err != nil {
				return fmt.Errorf("failed to encode manifest: %w", err)
			}

			path := fmt.Sprintf("%s/api/v1/modules", ctx.String("registry"))
			request, err := http.NewRequest("PUT", path, bytes.NewBuffer(bodyBz))
			if err != nil {
				return fmt.Errorf("failed to create request: %w", err)
			}

			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", credentials.Registry.Token))

			resp, err := client.Do(request)
			if err != nil {
				return fmt.Errorf("failed to publish module: %w", err)
			}

			defer func() {
				_ = resp.Body.Close()
			}()

			if resp.StatusCode != http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return fmt.Errorf("failed to read response body: %w", err)
				}

				return fmt.Errorf("failed to publish module: %w", errors.New(string(body)))
			}

			_, _ = color.New(color.FgGreen).Fprintln(ctx.App.Writer, "module successfully published!")
			return nil
		},
	}
}
