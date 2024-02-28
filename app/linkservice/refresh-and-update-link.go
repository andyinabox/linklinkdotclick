package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/feedfinder"
)

// todo: it makes more sense to me that the link should be saved here
func (s *Service) RefreshAndUpdateLink(userId uint, link app.Link, force bool) (*app.Link, error) {

	if link.FeedUrl == "" {
		s.log.Info().Printf("no feed found for %s, skipping refresh", link.SiteName)
		return s.UpdateLink(userId, link, false)
	}

	if force || time.Now().After(link.LastFetched.Add(s.conf.LinkRefreshBuffer)) {
		feedData, err := feedfinder.RefreshFeedDataForUrl(link.FeedUrl)
		if err != nil {
			// if refresh fails, see if we can update the feed url from the site
			feedData, err = feedfinder.GetFeedDataForUrl(link.SiteUrl)
			if err != nil {
				s.log.Warn().Printf("failed to refresh feed for %s, %v", link.SiteName, err)
				return &link, nil
			}
			link.FeedUrl = feedData.FeedUrl
		}

		// set unread count
		link.UnreadCount = int16(feedData.NewItemsCount(&link.LastClicked))
		link.LastFetched = time.Now()

	}

	return s.UpdateLink(userId, link, false)
}
