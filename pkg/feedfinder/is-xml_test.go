package feedfinder

import (
	"testing"

	"github.com/andyinabox/linkydink/pkg/responses"
	"github.com/andyinabox/linkydink/test"
)

func Test_IsXml(t *testing.T) {

	// test feed url
	{
		ts := test.NewFixtureTestServer("../../test/fixtures/www.w3c.org/feed.xml", t)
		body, err := responses.GetBodyFromUrl(ts.URL)
		if err != nil {
			t.Fatal(err.Error())
		}
		isXml := IsXml(body)
		if !isXml {
			t.Errorf("expected IsXml to be true for %s, got false", ts.URL)
		}
	}

	// test non-feed url
	{
		ts := test.NewFixtureTestServer("../../test/fixtures/www.w3c.org/index.html", t)
		body, err := responses.GetBodyFromUrl(ts.URL)
		if err != nil {
			t.Fatal(err.Error())
		}
		isXml := IsXml(body)
		if isXml {
			t.Errorf("expected IsXml to be false for %s, got false", ts.URL)
		}
	}
}
