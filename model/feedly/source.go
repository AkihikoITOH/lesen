package feedly

import (
	"github.com/AkihikoITOH/lesen/parser"
	"github.com/mmcdole/gofeed"
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

func (s *Source) Feed() *gofeed.Feed {
	return s.feed
}
