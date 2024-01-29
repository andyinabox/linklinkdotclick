package feedreader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/mmcdole/gofeed"
)

func (r *Reader) ParseFeed(rawurl string) (result *Result, err error) {
	var feedUrl *url.URL
	res, err := http.Get(rawurl)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	defer res.Body.Close()
	var b []byte
	b, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	docStart := strings.TrimSpace(string(b))[:5]

	// if no xml element is found, assume it's a site
	if docStart != "<?xml" {
		var urls []url.URL
		// TODO: figure out why getFeedUrlsFromResponse doesn't work here
		urls, err = r.GetFeedUrls(rawurl)
		if err != nil {
			return
		}
		if len(urls) == 0 {
			err = fmt.Errorf("no feeds found at %s", rawurl)
			return
		}
		firstUrl := urls[0]
		feedUrl = &firstUrl
	} else {
		feedUrl, err = url.Parse(rawurl)
	}

	fp := gofeed.NewParser()
	var feed *gofeed.Feed
	feed, err = fp.ParseURL(feedUrl.String())

	result = &Result{
		Feed: feed,
	}

	return
}
