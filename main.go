package main

import (
	"log"
	"os"

	"github.com/cosmos/atlas/cmd"
)

func main() {
	app := cmd.NewApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
