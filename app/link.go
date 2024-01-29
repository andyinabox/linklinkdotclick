package app

import "time"

type Link struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userId"`
	SiteName    string    `json:"siteName"`
	SiteUrl     string    `json:"siteUrl"`
	FeedUrl     string    `json:"feedUrl"`
	OriginalUrl string    `json:"originalUrl"`
	UnreadCount int16     `json:"unreadCount"`
	LastClicked time.Time `json:"lastClicked"`
	LastFetched time.Time `json:"lastFetched"`
}

type testData struct {
	Links []Link
}
