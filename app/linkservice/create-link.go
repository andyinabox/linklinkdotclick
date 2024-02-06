package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/util"
)

func newBaseLink(reqUrl string) *app.Link {
	return &app.Link{
		OriginalUrl: reqUrl,
		LastClicked: time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC),
		LastFetched: time.Now(),
	}
}

func (s *Service) createLinkFeedUrl(body []byte, reqUrl string) (*app.Link, error) {
	feedData, err := s.fs.ParseFeedResponse(body, reqUrl)
	if err != nil {
		return nil, err
	}
	link := newBaseLink(reqUrl)
	link.SiteName = feedData.SiteName()
	link.SiteUrl = feedData.SiteUrl()
	link.FeedUrl = reqUrl
	link.UnreadCount = int16(feedData.NewItemsCount(&link.LastClicked))

	return link, nil
}

func (s *Service) createLinkSiteUrl(body []byte, reqUrl string) (*app.Link, error) {

	siteData, err := s.fs.GetSiteData(body, reqUrl)
	if err != nil {
		return nil, err
	}
	feedUrls := siteData.FeedUrls()

	// no feeds found
	if len(feedUrls) == 0 {
		link := newBaseLink(reqUrl)
		link.SiteName = siteData.SiteName()
		link.SiteUrl = reqUrl
		return link, nil
	}

	body, err = util.GetResponseBodyFromUrl(feedUrls[0])
	if err != nil {
		return nil, err
	}
	return s.createLinkFeedUrl(body, reqUrl)
}

func (s *Service) CreateLink(url string) (*app.Link, error) {

	body, err := util.GetResponseBodyFromUrl(url)
	if err != nil {
		return nil, err
	}

	var link *app.Link
	if s.fs.IsXml(body) {
		link, err = s.createLinkFeedUrl(body, url)
	} else {
		link, err = s.createLinkSiteUrl(body, url)
	}
	if err != nil {
		return nil, err
	}

	return s.lr.CreateLink(*link)
}
