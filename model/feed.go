package model

import "time"

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
	Articles() []Article
}

type Article interface {
	Title() string
	Description() string
	Link() string
	PublishedAt() *time.Time
	MarkAsRead()
	MarkedAsRead() bool
	MarkAsUnread()
}
