package fixtures

import (
	"time"

	"github.com/andyinabox/linkydink/app"
)

func LinkBeforeInsertNoUserID() app.Link {
	now := time.Now()
	return app.Link{
		SiteName:    "W3C - Blog",
		SiteUrl:     "https://www.w3.org/blog/",
		FeedUrl:     "https://www.w3.org/blog/feed/",
		OriginalUrl: "https://www.w3.org/blog/",
		UnreadCount: 25,
		LastClicked: time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC),
		LastFetched: now,
	}
}

func LinkBeforeInsert() app.Link {
	link := LinkBeforeInsertNoUserID()
	link.UserID = 1
	return link
}

func LinkJustCreated() app.Link {
	link := LinkBeforeInsert()
	now := time.Now()
	link.ID = 1
	link.CreatedAt = now
	link.UpdatedAt = now
	link.LastFetched = now
	link.UnreadCount = 5
	return link
}

func LinkSetLastCLicked(t *time.Time) app.Link {
	link := LinkJustCreated()
	link.LastClicked = *t
	return link
}
