package feedservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/util"
	"github.com/andyinabox/linkydink/test"
)

func Test_GetSiteData(t *testing.T) {
	s := New()
	{
		ts := test.NewFixtureTestServer("../../test/fixtures/www.w3c.org/index.html", t)
		body, err := util.GetResponseBodyFromUrl(ts.URL)
		if err != nil {
			t.Fatal(err.Error())
		}

		data, err := s.GetSiteData(body, ts.URL)
		if err != nil {
			t.Fatal(err.Error())
		}

		expectedSiteName := "Blog | W3C"
		if data.SiteName() != expectedSiteName {
			t.Errorf("expected SiteName for %s to return %s, got %s", ts.URL, expectedSiteName, data.SiteName())
		}

		expectedFeedCount := 1
		expectedFirstFeedUrl := ts.URL + "/blog/feed/"
		feedUrls := data.FeedUrls()
		if len(feedUrls) != expectedFeedCount {
			t.Errorf("expected feed count for %s to be %d, got %d", ts.URL, expectedFeedCount, len(feedUrls))
		}
		if feedUrls[0] != expectedFirstFeedUrl {
			t.Errorf("expected feed url for %s to be %s, got %s", ts.URL, expectedFirstFeedUrl, feedUrls[0])
		}
	}
}
