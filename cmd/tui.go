package cmd

import (
	"github.com/urfave/cli"

	"github.com/AkihikoITOH/lesen/model/feedly"
	"github.com/AkihikoITOH/lesen/ui"
)

var tuiCommand = cli.Command{
	Name:    "tui",
	Aliases: []string{"t"},
	Usage:   "Open feed in TUI.",
	Action:  openTUIAction,
	Flags: []cli.Flag{
		inputFileLocationFlag,
		directoryTitlesFlag,
	},
}

func openTUIAction(c *cli.Context) error {
	root, err := feedly.NewFeedsFromOPML(c.String("input"))
	if err != nil {
		return err
	}

	root.FetchSources()

	tui := ui.NewTUI(root)
	tui.Draw()

	return nil
}
