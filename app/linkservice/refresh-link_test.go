package linkservice

import (
	"testing"
	"time"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/test"
)

func Test_RefreshLink(t *testing.T) {
	ts := test.NewFixtureTestServer("../../test/fixtures/www.w3c.org/feed.xml", t)
	ls := NewLinkService(t)

	dateFormat := "2006-Jan-02"
	afterDate, err := time.Parse(dateFormat, "2024-Jan-27")
	if err != nil {
		t.Fatal(err.Error())
	}

	link := app.Link{
		FeedUrl:     ts.URL + "/blog/feed",
		LastClicked: afterDate,
	}

	refreshed, err := ls.RefreshLink(link)
	if err != nil {
		t.Fatal(err.Error())
	}
	if refreshed.UnreadCount != 2 {
		t.Errorf("expected unread count to be 2, got %d", refreshed.UnreadCount)
	}
}
