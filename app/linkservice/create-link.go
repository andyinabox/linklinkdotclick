package linkservice

import (
	"errors"
	"net/url"
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

// seems like this should happen in feed service
func (s *Service) getSiteUrl(feedDataSiteUrl string, reqUrl string, siteData app.SiteData) (siteUrl string, err error) {
	siteUrl = feedDataSiteUrl
	if siteUrl == "" && siteData != nil {
		siteUrl = siteData.SiteUrl()
	}
	if siteUrl == "" {
		reqUrlData, _ := url.Parse(reqUrl)
		siteUrl = reqUrlData.Scheme + "://" + reqUrlData.Host
	}
	if siteUrl == "" {
		err = errors.New("could not parse site url")
	}
	return
}

// seems like this should happen in feed service
func (s *Service) getSiteName(feedDataSiteName string, siteUrl string, siteData app.SiteData) (siteName string, err error) {
	siteName = feedDataSiteName
	if siteName == "" {
		if siteData == nil && siteUrl != "" {
			var body []byte
			body, err = util.GetResponseBodyFromUrl(siteUrl)
			if err != nil {
				return
			}
			siteData, err = s.fs.GetSiteData(body, siteUrl)
		}
		if siteData != nil {
			siteName = siteData.SiteName()
		}
	}
	return
}

func (s *Service) createLinkFeedUrl(body []byte, reqUrl string, siteData app.SiteData) (*app.Link, error) {
	feedData, err := s.fs.ParseFeedResponse(body, reqUrl)
	if err != nil {
		return nil, err
	}
	link := newBaseLink(reqUrl)

	siteUrl, err := s.getSiteUrl(feedData.SiteUrl(), reqUrl, siteData)
	if err != nil {
		return nil, err
	}

	siteName, err := s.getSiteName(feedData.SiteName(), siteUrl, siteData)
	if err != nil {
		return nil, err
	}

	link.SiteName = siteName
	link.SiteUrl = siteUrl
	link.FeedUrl = feedData.FeedUrl()
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

	reqUrl = feedUrls[0]
	body, err = util.GetResponseBodyFromUrl(reqUrl)
	if err != nil {
		return nil, err
	}
	return s.createLinkFeedUrl(body, reqUrl, siteData)
}

func (s *Service) CreateLink(url string) (*app.Link, error) {

	body, err := util.GetResponseBodyFromUrl(url)
	if err != nil {
		return nil, err
	}

	var link *app.Link
	if s.fs.IsXml(body) {
		link, err = s.createLinkFeedUrl(body, url, nil)
	} else {
		link, err = s.createLinkSiteUrl(body, url)
	}
	if err != nil {
		return nil, err
	}

	return s.lr.CreateLink(*link)
}
