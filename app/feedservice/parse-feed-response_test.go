package feedservice

import (
	"testing"
	"time"

	"github.com/andyinabox/linkydink/app/util"
	"github.com/andyinabox/linkydink/test"
)

func Test_ParseFeed(t *testing.T) {
	ts := test.NewFixtureTestServer("../../test/fixtures/www.w3c.org/feed.xml", t)
	body, err := util.GetResponseBodyFromUrl(ts.URL)
	if err != nil {
		t.Fatal(err.Error())
	}

	s := New()

	d, err := s.ParseFeedResponse(body, ts.URL)
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedName := "W3C - Blog"
	expectedSiteUrl := "https://www.w3.org/blog/"
	expectedFeedUrl := "https://www.w3.org/blog/feed/"
	expectedNewItemsCount := 25

	// testing with date that should supply 2 items
	dateFormat := "2006-Jan-02"
	afterDate, err := time.Parse(dateFormat, "2024-Jan-27")
	if err != nil {
		t.Fatalf(err.Error())
	}
	expectedAfterDateCount := 2
	now := time.Now()

	if d.SiteName() != expectedName {
		t.Errorf("expected SiteName() to return %s, got %s", expectedName, d.SiteName())
	}
	if d.SiteUrl() != expectedSiteUrl {
		t.Errorf("expected SiteUrl() to return %s, got %s", expectedSiteUrl, d.SiteUrl())
	}
	if d.FeedUrl() != expectedFeedUrl {
		t.Errorf("expected FeedUrl() to return %s, got %s", expectedFeedUrl, d.FeedUrl())
	}
	if d.FeedUrl() != expectedFeedUrl {
		t.Errorf("expected FeedUrl() to return %s, got %s", expectedFeedUrl, d.FeedUrl())
	}
	if d.NewItemsCount(&now) != 0 {
		t.Errorf("expected %d items count for %v, got %d", 0, now, d.NewItemsCount(&now))
	}
	if d.NewItemsCount(&afterDate) != uint(expectedAfterDateCount) {
		t.Errorf("expected %d items count for %v, got %d", expectedAfterDateCount, afterDate, d.NewItemsCount(&afterDate))
	}
	if d.NewItemsCount(nil) != uint(expectedNewItemsCount) {
		t.Errorf("expected %d items count for nil date, got %d", expectedNewItemsCount, d.NewItemsCount(nil))
	}
}
