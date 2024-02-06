package linkservice

import (
	"net/http"
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) RefreshLink(link app.Link) (*app.Link, error) {

	// skip if there is no feed
	if link.FeedUrl == "" {
		return &link, nil
	}

	res, err := http.Get(link.FeedUrl)
	if err != nil {
		return nil, app.ErrServerError
	}

	if res.StatusCode != http.StatusOK {
		return nil, app.ErrNotFound
	}

	feedData, err := s.fs.ParseFeedResponse(res)
	if err != nil {
		return nil, err
	}

	// set unread count
	link.UnreadCount = int16(feedData.NewItemsCount(&link.LastClicked))
	link.LastFetched = time.Now()

	return &link, nil
}
