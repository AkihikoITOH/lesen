package model

import (
	"github.com/mmcdole/gofeed"
)

type Root interface {
	Title() string
	Directories() []Directory
	FetchSources()
}

type Directory interface {
	Title() string
	Sources() []Source
}

type Source interface {
	Title() string
	XMLURL() string
	HTMLURL() string
	Fetch() error
	Feed() *gofeed.Feed
}
