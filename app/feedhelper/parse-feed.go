package feedhelper

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/mmcdole/gofeed"
)

func (h *Helper) ParseFeed(feedUrl string) (app.FeedData, error) {

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedUrl)
	if err != nil {
		return nil, err
	}

	return &FeedData{*feed}, nil
}
