package test

import (
	"time"

	"github.com/andyinabox/linkydink/app"
)

func LinkBeforeInsert() *app.Link {
	now := time.Now()
	return &app.Link{
		SiteName:    "Feed for https://jamesg.blog/",
		SiteUrl:     "https://jamesg.blog/",
		FeedUrl:     "https://jamesg.blog/feeds/posts.xml",
		OriginalUrl: "https://jamesg.blog/",
		UnreadCount: 5,
		LastClicked: time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC),
		LastFetched: now,
	}
}

// func LinkJustClicked() *app.Link {
// 	link := LinkJustCreated()
// 	now := time.Now()
// 	link.ID = 1
// 	link.CreatedAt = now
// 	link.UpdatedAt = now
// 	link.LastClicked = now
// 	link.LastFetched = now
// 	link.UnreadCount = 0
// 	return link
// }
