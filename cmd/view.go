package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/AkihikoITOH/lesen/model/feedly"
	"github.com/AkihikoITOH/lesen/query"
)

var listCommand = cli.Command{
	Name:    "list",
	Aliases: []string{"ls"},
	Usage:   "List the feeds.",
	Action:  listAction,
	Flags: []cli.Flag{
		inputFileLocationFlag,
		directoryTitlesFlag,
		feedModeFlag,
	},
}

func listAction(c *cli.Context) error {
	root, err := feedly.NewFeedsFromOPML(c.String("input"))
	if err != nil {
		return err
	}
	root.FetchSources()

	q := &query.Query{Root: root}
	opts := &query.CollectDirectoriesOpts{Titles: c.StringSlice("titles")}
	directories := q.CollectDirectories(opts)

	for _, dir := range directories {
		fmt.Printf("* %s\n", dir.Title())
		for _, src := range dir.Sources() {
			feed := src.Feed()
			if feed == nil {
				fmt.Println(" * Feed unavailable.")
				continue
			}
			fmt.Printf(" * %s\n", src.Title())
			for _, item := range feed.Items {
				if len(item.Title) > 0 {
					fmt.Printf("  * %s\n", item.Title)
				}
			}
		}
	}

	return nil
}
