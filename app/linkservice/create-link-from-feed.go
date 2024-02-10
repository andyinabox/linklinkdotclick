package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/feedfinder"
)

func (s *Service) CreateLinkFromFeed(userId uint, siteTitle string, feedUrl string, siteUrl string) (*app.Link, error) {
	feedData, err := feedfinder.RefreshFeedDataForUrl(feedUrl)
	if err != nil {
		// if that doesn't work, try using siteUrl
		feedData, err = feedfinder.GetFeedDataForUrl(siteUrl)

		if err != nil {
			return nil, err
		}
	}

	lastClicked := defaultLastClickedDate()

	return s.lr.CreateLink(app.Link{
		UserID:      userId,
		OriginalUrl: feedUrl,
		LastClicked: lastClicked,
		LastFetched: time.Now(),
		SiteName:    siteTitle,
		SiteUrl:     feedData.SiteUrl,
		FeedUrl:     feedData.FeedUrl,
		UnreadCount: int16(feedData.NewItemsCount(&lastClicked)),
	})
}
