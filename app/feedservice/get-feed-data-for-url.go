package feedservice

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/util"
)

func (s *Service) GetFeedDataForUrl(originalUrl string) (app.FeedData, error) {
	body, err := util.GetResponseBodyFromUrl(originalUrl)
	if err != nil {
		return nil, err
	}

	var feedUrl = originalUrl

	if !s.IsXml(body) {
		siteData, err := s.GetSiteData(body, originalUrl)
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

	return s.refreshFeedDataForUrl(feedUrl, nil)
}
