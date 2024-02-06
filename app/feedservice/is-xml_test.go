package feedservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/util"
	"github.com/andyinabox/linkydink/test"
)

func Test_IsXml(t *testing.T) {
	s := New()

	// test feed url
	{
		ts := test.NewFixtureTestServer("../../test/fixtures/www.w3c.org/feed.xml", t)
		body, err := util.GetResponseBodyFromUrl(ts.URL)
		if err != nil {
			t.Fatal(err.Error())
		}
		isXml := s.IsXml(body)
		if !isXml {
			t.Errorf("expected IsXml to be true for %s, got false", ts.URL)
		}
	}

	// test non-feed url
	{
		ts := test.NewFixtureTestServer("../../test/fixtures/www.w3c.org/index.html", t)
		body, err := util.GetResponseBodyFromUrl(ts.URL)
		if err != nil {
			t.Fatal(err.Error())
		}
		isXml := s.IsXml(body)
		if isXml {
			t.Errorf("expected IsXml to be false for %s, got false", ts.URL)
		}
	}
}
