package feedservice

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/mmcdole/gofeed"
)

func (s *Service) ParseFeedResponse(body []byte, reqUrl string) (app.FeedData, error) {

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(body))
	if err != nil {
		return nil, err
	}

	feedData := &FeedData{*feed}

	if feedData.Feed.FeedLink == "" {
		feedData.Feed.FeedLink = reqUrl
	}

	return feedData, nil
}
