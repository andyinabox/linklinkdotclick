package feedservice

import (
	"errors"
	"net/url"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/util"
)

func (s *Service) refreshFeedDataForUrl(feedUrl string, siteData *SiteData) (app.FeedData, error) {
	body, err := util.GetResponseBodyFromUrl(feedUrl)
	if err != nil {
		return nil, err
	}
	feedData, err := s.ParseFeedResponse(body, feedUrl)
	if err != nil {
		return nil, err
	}

	// try differeny options for filling in missing site url
	if feedData.siteUrl == "" && siteData != nil {
		feedData.siteUrl = siteData.SiteUrl
	}
	if feedData.siteUrl == "" {
		reqUrlData, _ := url.Parse(feedUrl)
		feedData.siteUrl = reqUrlData.Scheme + "://" + reqUrlData.Host
	}
	if feedData.siteUrl == "" {
		return nil, errors.New("could not parse site url")
	}

	// try differeny options for filling in missing site name
	if feedData.siteName == "" {
		if siteData == nil && feedData.siteUrl != "" {
			body, err = util.GetResponseBodyFromUrl(feedData.siteUrl)
			if err != nil {
				return nil, err
			}
			siteData, err = s.GetSiteData(body, feedData.siteUrl)
			if err != nil {
				return nil, err
			}
		}
		if siteData != nil {
			feedData.siteName = siteData.SiteName
		}
	}
	if feedData.siteName == "" {
		return nil, errors.New("could not parse site name")
	}

	return feedData, nil
}

func (s *Service) RefreshFeedDataForUrl(feedUrl string) (app.FeedData, error) {
	return s.refreshFeedDataForUrl(feedUrl, nil)
}
