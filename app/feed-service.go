package app

import (
	"net/http"
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
	IsXml(*http.Response) bool
	GetSiteData(*http.Response) (SiteData, error)
	ParseFeedResponse(*http.Response) (FeedData, error)
}
