package cmd

import (
	"os"

	"github.com/urfave/cli"
)

func StartNewCLIApp() error {
	app := newCLIApp()
	return app.Run(os.Args)
}

func newCLIApp() *cli.App {
	app := cli.NewApp()
	app.Name = "lesen"
	app.Usage = "Command line RSS Reader"
	app.Commands = []cli.Command{listCommand}
	return app
}
