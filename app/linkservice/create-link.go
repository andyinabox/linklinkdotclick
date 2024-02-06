package linkservice

import (
	"net/http"
	"time"

	"github.com/andyinabox/linkydink/app"
)

func newBaseLink(res *http.Response) *app.Link {
	return &app.Link{
		OriginalUrl: res.Request.URL.String(),
		LastClicked: time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC),
		LastFetched: time.Now(),
	}
}

func (s *Service) createLinkFeedUrl(res *http.Response) (*app.Link, error) {
	feedData, err := s.fs.ParseFeedResponse(res)
	if err != nil {
		return nil, err
	}
	link := newBaseLink(res)
	link.SiteName = feedData.SiteName()
	link.SiteUrl = feedData.SiteUrl()
	link.FeedUrl = res.Request.URL.String()
	link.UnreadCount = int16(feedData.NewItemsCount(&link.LastClicked))

	return link, nil
}

func (s *Service) createLinkSiteUrl(res *http.Response) (*app.Link, error) {
	siteData, err := s.fs.GetSiteData(res)
	if err != nil {
		return nil, err
	}
	feedUrls := siteData.FeedUrls()

	// no feeds found
	if len(feedUrls) == 0 {
		link := newBaseLink(res)
		link.SiteName = siteData.SiteName()
		link.SiteUrl = res.Request.URL.String()
		return link, nil
	}

	feedRes, err := http.Get(feedUrls[0])
	if err != nil {
		return nil, app.ErrServerError
	}

	if feedRes.StatusCode != http.StatusOK {
		return nil, app.ErrNotFound
	}

	return s.createLinkFeedUrl(feedRes)
}

func (s *Service) CreateLink(url string) (*app.Link, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, app.ErrServerError
	}

	if res.StatusCode != http.StatusOK {
		return nil, app.ErrNotFound
	}

	var link *app.Link
	if s.fs.IsXml(res) {
		link, err = s.createLinkFeedUrl(res)
	} else {
		link, err = s.createLinkSiteUrl(res)
	}
	if err != nil {
		return nil, app.ErrServerError
	}

	return s.lr.CreateLink(*link)
}
