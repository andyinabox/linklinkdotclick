package linkservice

import (
	"errors"
	"fmt"
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) refreshLink(link app.Link) (*app.Link, error) {

	fmt.Println("Fetching rss feed for " + link.FeedUrl)
	result, err := s.fr.ParseFeedUrl(link.FeedUrl)
	if err != nil {
		return nil, err
	}
	if result.Feed == nil {
		return nil, errors.New("no rss feed found")
	}

	link.LastFetched = time.Now()

	// count feed items published after last click
	count := 0
	for _, item := range result.Items {
		if item.PublishedParsed.After(link.LastClicked) {
			count = count + 1
		}
	}

	// set unread count
	link.UnreadCount = int16(count)

	return &link, nil
}
