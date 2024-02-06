package app

import (
	"time"
)

type SiteData interface {
	SiteName() string
	FeedUrls() []string
}

type FeedData interface {
	SiteName() string
	SiteUrl() string
	FeedUrl() string
	NewItemsCount(after *time.Time) uint
}

type FeedService interface {
	IsXml(body []byte) bool
	GetSiteData(body []byte, reqUrl string) (SiteData, error)
	ParseFeedResponse(body []byte, reqUrl string) (FeedData, error)
}
