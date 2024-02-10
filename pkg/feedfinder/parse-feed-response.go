package feedfinder

import (
	"github.com/mmcdole/gofeed"
)

func ParseFeedResponse(body []byte, reqUrl string) (*FeedData, error) {

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(body))
	if err != nil {
		return nil, err
	}

	feedData := NewFeedDataFromFeed(feed)

	if feedData.FeedUrl == "" {
		feedData.FeedUrl = reqUrl
	}

	return feedData, nil
}
