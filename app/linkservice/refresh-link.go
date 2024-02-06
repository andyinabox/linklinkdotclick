package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) RefreshLink(link app.Link) (*app.Link, error) {

	// skip if there is no feed
	if link.FeedUrl == "" {
		return &link, nil
	}

	feedData, err := s.fs.ParseFeed(link.FeedUrl)
	if err != nil {
		return nil, err
	}

	// set unread count
	link.UnreadCount = int16(feedData.NewItemsCount(&link.LastClicked))
	link.LastFetched = time.Now()

	return &link, nil
}
