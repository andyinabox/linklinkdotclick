package feedreader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// TODO: this is making 2-3 requests when it could make just 1
func (r *Reader) Parse(rawurl string) (result *Result, err error) {
	feedUrl := rawurl
	res, err := http.Get(feedUrl)
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
		feedUrl = firstUrl.String()
	}

	return r.ParseFeedUrl(feedUrl)
}
