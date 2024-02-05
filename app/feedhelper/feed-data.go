package feedhelper

import (
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedData struct {
	gofeed.Feed
}

func (d *FeedData) SiteName() string {
	return strings.TrimSpace(d.Title)
}

func (d *FeedData) SiteUrl() string {
	return strings.TrimSpace(d.Link)
}

func (d *FeedData) FeedUrl() string {
	return strings.TrimSpace(d.FeedLink)
}

func (d *FeedData) NewItemsCount(after *time.Time) uint {
	if after == nil {
		return uint(len(d.Items))
	}

	var count uint = 0
	for _, item := range d.Items {
		if item.PublishedParsed.After(*after) {
			count += 1
		}
	}

	return count
}
