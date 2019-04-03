package cmd

import (
	"github.com/urfave/cli"
)

var (
	inputFileLocationFlag = cli.StringFlag{
		Name:   "input",
		Usage:  "Path of the desired input OPML file.",
		EnvVar: "LESEN_INPUT_FILE_LOCATION",
	}

	directoryTitlesFlag = cli.StringSliceFlag{
		Name:   "titles",
		Value:  &cli.StringSlice{},
		Usage:  "Comma-separated list of titles to show.",
		EnvVar: "LESEN_TITLES",
	}

	feedModeFlag = cli.StringFlag{
		Name:   "mode",
		Value:  "feedly",
		Usage:  "Feed mode",
		EnvVar: "LESEN_FEED_MODE",
	}

	debugFlag = cli.BoolFlag{
		Name:   "debug",
		Usage:  "Run in debug mode",
		EnvVar: "LESEN_DEBUG_MODE",
	}
)
