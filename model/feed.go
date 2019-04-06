package model

import (
	"github.com/mmcdole/gofeed"
)

type Root interface {
	Title() string
	SetTitle(string)
	Directories() []Directory
	SetDirectories([]Directory)
	FetchSources()
	Duplicate() Root
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
