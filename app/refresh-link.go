package app

import (
	"fmt"
	"time"
)

func (a *App) RefreshLink(link *Link) error {
	fmt.Printf("Fetching rss feed for %s\n", link.FeedUrl)
	result, err := a.feedreader.ParseFeedUrl(link.FeedUrl)
	if err != nil {
		return err
	}
	if result.Feed != nil {
		link.LastFetched = time.Now()

		// count feed items published after last click
		count := 0
		for _, item := range result.Items {
			if item.PublishedParsed.After(link.LastClicked) {
				count = count + 1
			}
		}

		// set unread count
		link.UnreadCount = int16(count)

		// save updated result
		tx := a.db.Save(link)
		err = tx.Error
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("No rss results for %s\n", link.FeedUrl)
	}
	return nil
}
