package linkservice

import (
	"io/ioutil"
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
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, app.ErrServerError
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
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
