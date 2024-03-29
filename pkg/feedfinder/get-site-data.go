package feedfinder

import (
	"bytes"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SiteData struct {
	SiteName string
	SiteUrl  string
	FeedUrls []string
}

func GetSiteData(body []byte, reqUrl string) (data *SiteData, err error) {

	bodyReader := ioutil.NopCloser(bytes.NewBuffer(body))
	doc, err := goquery.NewDocumentFromReader(bodyReader)
	if err != nil {
		return
	}

	feedUrls := []string{}
	title := doc.Find("title").First().Text()

	reqUrlData, err := url.Parse(reqUrl)
	if err != nil {
		return
	}

	doc.Find("[rel='alternate'][type^='application']").
		Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title

			feedType, _ := s.Attr("type")
			if !strings.Contains(feedType, "xml") {
				return
			}

			href, _ := s.Attr("href")
			if href != "" {
				var feedUrl *url.URL
				feedUrl, err = url.Parse(href)
				if err != nil {
					return
				}
				if feedUrl.Host == "" {
					feedUrl, err = url.Parse(reqUrlData.Scheme + "://" + reqUrlData.Host + feedUrl.String())
					if err != nil {
						return
					}
				}

				feedUrls = append(feedUrls, feedUrl.String())
			}
		})

	return &SiteData{
		SiteName: title,
		FeedUrls: feedUrls,
		SiteUrl:  reqUrl,
	}, nil
}
