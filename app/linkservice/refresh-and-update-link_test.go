package linkservice

import (
	"testing"
	"time"

	"github.com/andyinabox/linkydink/test"
)

func Test_RefreshAndUpdateLink(t *testing.T) {
	ts := test.NewFixtureTestServer("../../test/fixtures/www.w3c.org/feed.xml", t)
	ls := NewLinkService(t)

	dateFormat := "2006-Jan-02"
	afterDate, err := time.Parse(dateFormat, "2024-Jan-27")
	if err != nil {
		t.Fatal(err.Error())
	}

	link, err := ls.CreateLink(1, ts.URL+"/blog/feed")
	if err != nil {
		t.Fatal(err.Error())
	}
	link.LastClicked = afterDate

	refreshed, err := ls.RefreshAndUpdateLink(1, *link, true)
	if err != nil {
		t.Fatal(err.Error())
	}
	if refreshed.UnreadCount != 2 {
		t.Errorf("expected unread count to be 2, got %d", refreshed.UnreadCount)
	}
}
