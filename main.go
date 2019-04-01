package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/AkihikoITOH/lesen/cmd"
	"github.com/AkihikoITOH/lesen/model/feedly"
)

func main() {
	err := cmd.StartNewCLIApp()

	if err != nil {
		logrus.Fatal(err.Error())
	}
}

func test() {
	root, err := feedly.NewFeedsFromOPML("./local/data/small.opml")
	root.FetchSources()

	if err != nil {
		fmt.Printf(err.Error())
	} else {
		for _, dir := range root.Directories() {
			fmt.Printf("* %s\n", dir.Title())
			for _, src := range dir.Sources() {
				feed := src.Feed()
				if feed == nil {
					fmt.Println("  * Feed unavailable.")
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
	}
}
