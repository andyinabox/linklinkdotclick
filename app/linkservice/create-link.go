package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/feedfinder"
)

func (s *Service) CreateLink(userId uint, originalUrl string) (*app.Link, error) {

	feedData, err := feedfinder.GetFeedDataForUrl(originalUrl)
	if err != nil {
		return nil, err
	}

	lastClicked := defaultLastClickedDate()

	return s.lr.CreateLink(app.Link{
		UserID:      userId,
		OriginalUrl: originalUrl,
		LastClicked: lastClicked,
		LastFetched: time.Now(),
		SiteName:    feedData.SiteName,
		SiteUrl:     feedData.SiteUrl,
		FeedUrl:     feedData.FeedUrl,
		UnreadCount: int16(feedData.NewItemsCount(&lastClicked)),
	})
}
