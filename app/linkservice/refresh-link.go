package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) RefreshLink(userId uint, link app.Link) (*app.Link, error) {

	if link.FeedUrl == "" {
		s.log.Info().Printf("no feed found for %s, skipping refresh", link.SiteName)
		return &link, nil
	}

	feedData, err := s.fs.RefreshFeedDataForUrl(link.FeedUrl)
	if err != nil {
		// if refresh fails, see if we can update the feed url from the site
		feedData, err = s.fs.GetFeedDataForUrl(link.SiteUrl)
		if err != nil {
			s.log.Warn().Printf("failed to refresh feed for %s, %v", link.SiteName, err)
			return &link, nil
		}
		link.FeedUrl = feedData.FeedUrl()
	}

	// set unread count
	link.UnreadCount = int16(feedData.NewItemsCount(&link.LastClicked))
	link.LastFetched = time.Now()

	return &link, nil
}
