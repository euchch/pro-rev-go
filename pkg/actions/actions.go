package actions

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func ListItems(context *cli.Context) error {
	log.Infof("Displaying all items here")
	return nil
}
func CommandNotFoundAction(context *cli.Context, command string) {
	log.Fatalf("Command not found: `%v`\nPlease run `help` for list of commands\n", command)
}
