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
	Subcommands: []cli.Command{
		listDirectoriesCommand,
		listSourcesCommand,
		listArticlesCommand,
	},
}

var listDirectoriesCommand = cli.Command{
	Name:    "directories",
	Aliases: []string{"d"},
	Usage:   "List the directories.",
	Action:  listDirectoriesAction,
	Flags: []cli.Flag{
		inputFileLocationFlag,
		feedModeFlag,
	},
}

var listSourcesCommand = cli.Command{
	Name:    "sources",
	Aliases: []string{"s"},
	Usage:   "List the feed sources.",
	Action:  listSourcesAction,
	Flags: []cli.Flag{
		inputFileLocationFlag,
		directoryTitlesFlag,
		feedModeFlag,
	},
}

var listArticlesCommand = cli.Command{
	Name:    "articles",
	Aliases: []string{"a"},
	Usage:   "List the articles.",
	Action:  listArticlesAction,
	Flags: []cli.Flag{
		inputFileLocationFlag,
		directoryTitlesFlag,
		feedModeFlag,
	},
}

func listDirectoriesAction(c *cli.Context) error {
	root, err := feedly.NewFeedsFromOPML(c.String("input"))
	if err != nil {
		return err
	}

	for _, dir := range root.Directories() {
		fmt.Printf("* %s\n", dir.Title())
	}

	return nil
}

func listSourcesAction(c *cli.Context) error {
	root, err := feedly.NewFeedsFromOPML(c.String("input"))
	if err != nil {
		return err
	}

	q := &query.Query{Root: root}
	opts := query.CollectDirectoriesOpts{Titles: c.StringSlice("directories")}
	q = q.CollectDirectories(opts)

	filtered := q.Root

	for _, dir := range filtered.Directories() {
		fmt.Printf("* %s\n", dir.Title())
		for _, src := range dir.Sources() {
			fmt.Printf(" * %s\n", src.Title())
		}
	}

	return nil
}

func listArticlesAction(c *cli.Context) error {
	root, err := feedly.NewFeedsFromOPML(c.String("input"))
	if err != nil {
		return err
	}

	q := &query.Query{Root: root}
	opts := query.CollectDirectoriesOpts{Titles: c.StringSlice("directories")}
	q = q.CollectDirectories(opts)

	filtered := q.Root
	filtered.FetchSources()

	for _, dir := range filtered.Directories() {
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
					fmt.Printf("    %s\n", item.Link)
				}
			}
		}
	}

	return nil
}
