package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) CreateLink(userId uint, originalUrl string) (*app.Link, error) {
	feedData, err := s.fs.GetFeedDataForUrl(originalUrl)
	if err != nil {
		return nil, err
	}

	lastClicked := time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC)

	return s.lr.CreateLink(app.Link{
		UserID:      userId,
		OriginalUrl: originalUrl,
		LastClicked: lastClicked,
		LastFetched: time.Now(),
		SiteName:    feedData.SiteName(),
		SiteUrl:     feedData.SiteUrl(),
		FeedUrl:     feedData.FeedUrl(),
		UnreadCount: int16(feedData.NewItemsCount(&lastClicked)),
	})
}
