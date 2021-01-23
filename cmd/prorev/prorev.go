package main

import (
	"log"
	"os"

	"github.com/euchch/pro-rev-go/internal/proRevActions"
	"github.com/urfave/cli"
)

func runApp(app *cli.App) error {
	return app.Run(os.Args)
}

// imports as package "cli"
func main() {
	proRevApp := proRevActions.CreateApp()

	err := proRevApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
