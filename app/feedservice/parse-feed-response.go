package feedservice

import (
	"github.com/mmcdole/gofeed"
)

func (s *Service) ParseFeedResponse(body []byte, reqUrl string) (*FeedData, error) {

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(body))
	if err != nil {
		return nil, err
	}

	feedData := NewFeedDataFromFeed(feed)

	if feedData.feedUrl == "" {
		feedData.feedUrl = reqUrl
	}

	return feedData, nil
}
