package feedfinder

import (
	"errors"
	"net/url"

	"github.com/andyinabox/linkydink/pkg/responses"
)

func refreshFeedDataForUrl(feedUrl string, siteData *SiteData) (*FeedData, error) {
	body, err := responses.GetBodyFromUrl(feedUrl)
	if err != nil {
		return nil, err
	}
	feedData, err := ParseFeedResponse(body, feedUrl)
	if err != nil {
		return nil, err
	}

	// try differeny options for filling in missing site url
	if feedData.SiteUrl == "" && siteData != nil {
		feedData.SiteUrl = siteData.SiteUrl
	}
	if feedData.SiteUrl == "" {
		reqUrlData, _ := url.Parse(feedUrl)
		feedData.SiteUrl = reqUrlData.Scheme + "://" + reqUrlData.Host
	}
	if feedData.SiteUrl == "" {
		return nil, errors.New("could not parse site url")
	}

	// try differeny options for filling in missing site name
	if feedData.SiteName == "" {
		if siteData == nil && feedData.SiteUrl != "" {
			body, err = responses.GetBodyFromUrl(feedData.SiteUrl)
			if err != nil {
				return nil, err
			}
			siteData, err = GetSiteData(body, feedData.SiteUrl)
			if err != nil {
				return nil, err
			}
		}
		if siteData != nil {
			feedData.SiteName = siteData.SiteName
		}
	}
	if feedData.SiteName == "" {
		return nil, errors.New("could not parse site name")
	}

	return feedData, nil
}

func RefreshFeedDataForUrl(feedUrl string) (*FeedData, error) {
	return refreshFeedDataForUrl(feedUrl, nil)
}
