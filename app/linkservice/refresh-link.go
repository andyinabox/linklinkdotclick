package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/util"
)

func (s *Service) RefreshLink(link app.Link) (*app.Link, error) {

	// skip if there is no feed
	if link.FeedUrl == "" {
		return &link, nil
	}

	body, err := util.GetResponseBodyFromUrl(link.FeedUrl)
	if err != nil {
		return nil, err
	}

	feedData, err := s.fs.ParseFeedResponse(body, link.FeedUrl)
	if err != nil {
		return nil, err
	}

	// set unread count
	link.UnreadCount = int16(feedData.NewItemsCount(&link.LastClicked))
	link.LastFetched = time.Now()

	return &link, nil
}
