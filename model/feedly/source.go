package feedly

import (
	"github.com/AkihikoITOH/lesen/model"
	"github.com/AkihikoITOH/lesen/parser"
	"github.com/mmcdole/gofeed"
	"github.com/thoas/go-funk"
)

type Source struct {
	title   string
	xmlURL  string
	htmlURL string
	feed    *gofeed.Feed
}

func (s *Source) Title() string {
	return s.title
}

func (s *Source) XMLURL() string {
	return s.xmlURL
}

func (s *Source) HTMLURL() string {
	return s.htmlURL
}

func (s *Source) Fetch() error {
	feed, err := parser.LoadRSS(s.xmlURL)
	s.feed = feed
	return err
}

func (s *Source) Articles() []model.Article {
	if s.feed == nil {
		return []model.Article{}
	}

	articles, _ := funk.Map(
		s.feed.Items,
		func(item *gofeed.Item) model.Article {
			return &Article{feedItem: item}
		},
	).([]model.Article)

	return articles
}
