package feedreader

import "github.com/mmcdole/gofeed"

func (r *Reader) ParseFeedUrl(rawurl string) (result *Result, err error) {
	result = &Result{}

	fp := gofeed.NewParser()
	var feed *gofeed.Feed
	feed, err = fp.ParseURL(rawurl)

	result.Feed = feed

	return
}
