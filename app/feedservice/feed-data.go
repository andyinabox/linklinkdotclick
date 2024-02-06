package feedservice

import (
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedData struct {
	gofeed.Feed
	siteName string
	siteUrl  string
	feedUrl  string
	items    []*gofeed.Item
}

func NewFeedDataFromFeed(feed *gofeed.Feed) *FeedData {
	return &FeedData{
		Feed:     *feed,
		siteName: strings.TrimSpace(feed.Title),
		siteUrl:  strings.TrimSpace(feed.Link),
		feedUrl:  strings.TrimSpace(feed.FeedLink),
	}
}
func NewFeedDataFromSiteData(siteData *SiteData) *FeedData {
	return &FeedData{
		siteName: siteData.SiteName,
		siteUrl:  siteData.SiteUrl,
	}
}

func (d *FeedData) SiteName() string {
	return d.siteName
}

func (d *FeedData) SiteUrl() string {
	return d.siteUrl
}

func (d *FeedData) FeedUrl() string {
	return d.feedUrl
}

func (d *FeedData) NewItemsCount(after *time.Time) uint {
	if d.feedUrl == "" {
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
