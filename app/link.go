package app

import "time"

type Link struct {
	ID          int64
	UserID      int64
	SiteName    string
	SiteUrl     string
	FeedUrl     string
	OriginalUrl string
	UnreadCount int16
	LastClicked time.Time
	LastFetched time.Time
}
