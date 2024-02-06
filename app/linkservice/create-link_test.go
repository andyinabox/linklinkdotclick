package linkservice

import (
	"testing"

	"github.com/andyinabox/linkydink/test"
)

func Test_CreateLink(t *testing.T) {
	ts := test.NewFixtureTestServer("../../test/fixtures/www.w3c.org/feed.xml", t)
	ls := NewLinkService(t)
	_, err := ls.CreateLink(1, ts.URL)
	if err != nil {
		t.Fatal(err.Error())
	}
}
