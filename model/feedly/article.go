package feedly

import (
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/AkihikoITOH/lesen/config"
)

type Article struct {
	feedItem *gofeed.Item
}

func (a *Article) Title() string {
	return a.feedItem.Title
}

func (a *Article) Description() string {
	return a.feedItem.Description
}

func (a *Article) Link() string {
	return a.feedItem.Link
}

func (a *Article) PublishedAt() *time.Time {
	return a.feedItem.PublishedParsed
}

func (a *Article) MarkAsRead() {
	config.Backend().Write(a.Title())
}

func (a *Article) MarkedAsRead() bool {
	return config.Backend().Read(a.Title())
}

func (a *Article) MarkAsUnread() {
	config.Backend().Erase(a.Title())
}
