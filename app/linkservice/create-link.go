package linkservice

import (
	"errors"
	"strings"
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) CreateLink(url string) (*app.Link, error) {

	feed, feedUrl, err := s.fr.Parse(url)
	if err != nil {
		return nil, err
	}
	if feed == nil {
		return nil, errors.New("no feed detected")
	}

	link := &app.Link{
		SiteName:    strings.TrimSpace(feed.Title),
		SiteUrl:     strings.TrimSpace(feed.Link),
		FeedUrl:     feedUrl,
		OriginalUrl: url,
		UnreadCount: int16(len(feed.Items)),
		LastClicked: time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC),
		LastFetched: time.Now(),
	}

	return s.lr.CreateLink(*link)
}
