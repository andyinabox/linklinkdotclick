package feedservice

import (
	"testing"

	"github.com/andyinabox/linkydink/test"
)

func Test_IsXml(t *testing.T) {
	s := New()

	// test feed url
	{
		feedUrl := "https://www.w3.org/blog/feed/"
		r := test.NewMockResponse(feedUrl, "../../test/fixtures/www.w3.org/feed.xml", t)
		isXml := s.IsXml(r)
		if !isXml {
			t.Errorf("expected IsFeed to be true for %s, got false", feedUrl)
		}
	}

	// test non-feed url
	{
		siteUrl := "https://www.w3.org/blog/"
		r := test.NewMockResponse(siteUrl, "../../test/fixtures/www.w3.org/index.html", t)
		isXml := s.IsXml(r)
		if isXml {
			t.Errorf("expected IsFeed to be false for %s, got false", siteUrl)
		}
	}
}
