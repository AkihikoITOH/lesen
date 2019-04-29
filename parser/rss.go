package parser

import (
	"net/http"
	"time"

	"github.com/AkihikoITOH/lesen/config"
	"github.com/mmcdole/gofeed"
)

func LoadRSS(url string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	fp.Client = &http.Client{Timeout: time.Second * 5}
	config.Logger().Infof("Fetching %s", url)
	feed, err := fp.ParseURL(url)
	if err != nil {
		config.Logger().Warnf("Error while fetching %s (%s)", url, err.Error())
		return nil, err
	}
	config.Logger().Infof("Successfully fetched %s", url)
	return feed, nil
}
