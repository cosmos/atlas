package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"

	"github.com/cosmos/atlas/cmd"
)

func main() {
	// Load any environment variables found in the root .env if it exists. Note,
	// any environment variables manually provided will take precedence over
	// environment variables found in the .env file.
	_ = godotenv.Load()

	app := cmd.NewApp()

	if err := app.Run(os.Args); err != nil {
		_, _ = color.New(color.FgRed).Fprintln(app.ErrWriter, err.Error())
		os.Exit(1)
	}
}
