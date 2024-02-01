package feedreader

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func (r *Reader) GetFeedUrls(rawurl string) (feedUrls []url.URL, err error) {

	res, err := http.Get(rawurl)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	return getFeedUrlsFromResponse(res)

}

func getFeedUrlsFromResponse(res *http.Response) (feedUrls []url.URL, err error) {
	defer res.Body.Close()

	siteUrl := res.Request.URL

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	doc.Find("[rel='alternate'][type='application/rss+xml']").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		href, _ := s.Attr("href")
		if href != "" {
			var feedUrl *url.URL
			feedUrl, err = url.Parse(href)
			if err != nil {
				return
			}
			if feedUrl.Host == "" {
				feedUrl, err = url.Parse(siteUrl.Scheme + "://" + siteUrl.Host + feedUrl.String())
				if err != nil {
					return
				}
			}

			feedUrls = append(feedUrls, *feedUrl)
		}
	})

	return

}
