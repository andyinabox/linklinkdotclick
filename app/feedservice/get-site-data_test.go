package feedservice

import (
	"net/http"
	"testing"
)

func Test_GetSiteData(t *testing.T) {
	s := New()
	{
		url := "https://www.w3.org/blog/"
		r, err := http.Get(url)
		if err != nil {
			t.Fatal(err.Error())
		}
		data, err := s.GetSiteData(r)
		if err != nil {
			t.Fatal(err.Error())
		}

		expectedSiteName := "Blog | W3C"
		if data.SiteName() != expectedSiteName {
			t.Errorf("expected SiteName for %s to return %s, got %s", url, expectedSiteName, data.SiteName())
		}

		expectedFeedCount := 1
		expectedFirstFeedUrl := "https://www.w3.org/blog/feed/"
		feedUrls := data.FeedUrls()
		if len(feedUrls) != expectedFeedCount {
			t.Errorf("expected feed count for %s to be %d, got %d", url, expectedFeedCount, len(feedUrls))
		}
		if feedUrls[0] != expectedFirstFeedUrl {
			t.Errorf("expected feed url for %s to be %s, got %s", url, expectedFirstFeedUrl, feedUrls[0])
		}
	}
}
