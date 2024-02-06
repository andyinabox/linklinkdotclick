package feedservice

import (
	"io"
	"net/http"

	"github.com/andyinabox/linkydink/app"
	"github.com/mmcdole/gofeed"
)

func (s *Service) ParseFeedResponse(res *http.Response) (app.FeedData, error) {

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(b))
	if err != nil {
		return nil, err
	}

	feedData := &FeedData{*feed}

	if feedData.Feed.FeedLink == "" {
		feedData.Feed.FeedLink = res.Request.URL.String()
	}

	return feedData, nil
}
