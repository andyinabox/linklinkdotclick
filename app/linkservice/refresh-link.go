package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) RefreshLink(userId uint, link app.Link) (*app.Link, error) {

	if link.FeedUrl == "" {
		return &link, nil
	}

	feedData, err := s.fs.RefreshFeedDataForUrl(link.FeedUrl)
	if err != nil {
		return nil, err
	}

	// set unread count
	link.UnreadCount = int16(feedData.NewItemsCount(&link.LastClicked))
	link.LastFetched = time.Now()

	return &link, nil
}
