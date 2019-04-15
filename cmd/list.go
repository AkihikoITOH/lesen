package cmd

import (
	"github.com/urfave/cli"

	"github.com/AkihikoITOH/lesen/model/feedly"
	"github.com/AkihikoITOH/lesen/query"
	"github.com/AkihikoITOH/lesen/ui"
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

	cui := ui.NewCUI(root, ui.DirectoryLevel)
	cui.Draw()

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

	cui := ui.NewCUI(filtered, ui.SourceLevel)
	cui.Draw()

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

	cui := ui.NewCUI(filtered, ui.ArticleLevel)
	cui.Draw()

	return nil
}
