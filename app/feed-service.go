package app

import (
	"time"
)

type FeedData interface {
	SiteName() string
	SiteUrl() string
	FeedUrl() string
	NewItemsCount(after *time.Time) uint
}

type FeedService interface {
	GetFeedDataForUrl(string) (FeedData, error)
	RefreshFeedDataForUrl(string) (FeedData, error)
}
