package linkservice

import (
	"net/http"
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) CreateLink(url string) (*app.Link, error) {

	link := &app.Link{
		OriginalUrl: url,
		LastClicked: time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC),
		LastFetched: time.Now(),
	}

	feedUrl := url
	var siteData app.SiteData

	res, err := http.Get(url)
	if err != nil {
		return nil, app.ErrServerError
	}

	if res.StatusCode != http.StatusOK {
		return nil, app.ErrNotFound
	}

	if !s.fh.IsFeed(res) {
		siteData, err = s.fh.GetSiteData(res)
		if err != nil {
			return nil, err
		}
		feedUrls := siteData.FeedUrls()

		// no feeds found
		if len(feedUrls) == 0 {
			link.FeedFound = false
			link.SiteName = siteData.SiteName()
			link.SiteUrl = url
			return link, nil
		}

		feedUrl = feedUrls[0]
	}

	feedData, err := s.fh.ParseFeed(feedUrl)
	if err != nil {
		return nil, err
	}

	link.FeedFound = true
	link.SiteName = feedData.SiteName()
	link.SiteUrl = feedData.SiteUrl()
	link.FeedUrl = feedUrl
	link.UnreadCount = int16(feedData.NewItemsCount(&link.LastClicked))

	return s.lr.CreateLink(*link)
}
