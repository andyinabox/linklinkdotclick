package feedhelper

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/andyinabox/linkydink/app"
)

type SiteData struct {
	siteName string
	feedUrls []string
}

func (s *SiteData) SiteName() string {
	return s.siteName
}

func (s *SiteData) FeedUrls() []string {
	return s.feedUrls
}

func (h *Helper) GetSiteData(res *http.Response) (data app.SiteData, err error) {

	siteUrl := res.Request.URL

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	feedUrls := []string{}
	title := doc.Find("title").First().Text()

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
					feedUrl, err = url.Parse(siteUrl.Scheme + "://" + siteUrl.Host + feedUrl.String())
					if err != nil {
						return
					}
				}

				feedUrls = append(feedUrls, feedUrl.String())
			}
		})

	return &SiteData{
		siteName: title,
		feedUrls: feedUrls,
	}, nil
}
