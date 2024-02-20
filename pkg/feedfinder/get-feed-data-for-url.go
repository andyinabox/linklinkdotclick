package feedfinder

import (
	"net/url"

	"github.com/andyinabox/linkydink/pkg/responses"
)

func GetFeedDataForUrl(originalUrl string) (*FeedData, error) {

	parsedUrl, err := url.Parse(originalUrl)
	if err != nil {
		return nil, err
	}

	// if scheme is empty, try http
	if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "http"
	}

	originalUrl = parsedUrl.String()

	body, err := responses.GetBodyFromUrl(originalUrl)
	if err != nil {
		return nil, err
	}

	var feedUrl = originalUrl

	if !IsXml(body) {
		siteData, err := GetSiteData(body, originalUrl)
		if err != nil {
			return nil, err
		}

		// non-feed link
		if len(siteData.FeedUrls) == 0 {
			feedData := NewFeedDataFromSiteData(siteData)
			return feedData, nil
		}

		// site link that has a feed, get first feed
		feedUrl = siteData.FeedUrls[0]
	}

	return refreshFeedDataForUrl(feedUrl, nil)
}
