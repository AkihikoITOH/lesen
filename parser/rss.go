package parser

import (
	"fmt"

	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
)

func LoadRSS(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	logrus.Infof("Fetching %s", url)
	feed, err := fp.ParseURL(url)
	if err != nil {
		logrus.Warn(fmt.Sprintf("Error while fetching %s (%s)", url, err.Error()))
		return nil, err
	}
	logrus.Infof("Successfully fetched %s", url)
	return feed, nil
}
