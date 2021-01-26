package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	// Version defines the binary version at compile-time
	Version = ""

	// Commit defines the binary commit hash at compile-time
	Commit = ""

	// Branch defines the binary branch at compile-time
	Branch = ""
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Atlas CLI"
	app.Usage = "A Cosmos SDK module registry framework"
	app.Version = getVersion()
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	app.Commands = []*cli.Command{
		StartServerCommand(),
		LoginCommand(),
		PublishCommand(),
		InitCommand(),
	}

	return app
}

func getVersion() string {
	switch {
	case Version != "" && Commit != "":
		return fmt.Sprintf("%s-%s", Version, Commit)

	case Version != "":
		return Version

	case Branch != "" && Commit != "":
		return fmt.Sprintf("%s-%s", Branch, Commit)

	case Commit != "":
		return Commit

	case Branch != "":
		return Branch

	default:
		return ""
	}
}
