package proRevActions

import (
	"github.com/euchch/pro-rev-go/pkg/actions"
	"github.com/urfave/cli"
)

func CreateApp() *cli.App {
	return &cli.App{
		Name:  "pro-rev",
		Usage: "Professional reviewer management tool",
		Commands: []cli.Command{
			{
				Name:   "list",
				Usage:  "List all existing items",
				Action: actions.ListItems,
			},
		},
		CommandNotFound: actions.CommandNotFoundAction,
	}
}
