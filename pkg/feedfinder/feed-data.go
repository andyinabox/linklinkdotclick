package feedfinder

import (
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedData struct {
	gofeed.Feed
	SiteName string
	SiteUrl  string
	FeedUrl  string
}

func NewFeedDataFromFeed(feed *gofeed.Feed) *FeedData {
	return &FeedData{
		Feed:     *feed,
		SiteName: strings.TrimSpace(feed.Title),
		SiteUrl:  strings.TrimSpace(feed.Link),
		FeedUrl:  strings.TrimSpace(feed.FeedLink),
	}
}
func NewFeedDataFromSiteData(siteData *SiteData) *FeedData {
	return &FeedData{
		SiteName: siteData.SiteName,
		SiteUrl:  siteData.SiteUrl,
	}
}

func (d *FeedData) NewItemsCount(after *time.Time) uint {
	if d.FeedUrl == "" {
		return 0
	}

	if after == nil {
		return uint(len(d.Items))
	}

	var count uint = 0
	for _, item := range d.Items {
		if item.PublishedParsed.After(*after) {
			count = count + 1
		}
	}

	return count
}
