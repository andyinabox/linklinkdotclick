package linkservice

import (
	"testing"

	"github.com/andyinabox/linkydink/test/fixtures"
)

func Test_RefreshLink(t *testing.T) {
	ls := NewLinkService(t)

	link := fixtures.LinkJustClicked()
	refreshed, err := ls.RefreshLink(link)
	if err != nil {
		t.Fatal(err.Error())
	}
	if refreshed.UnreadCount != 0 {
		t.Errorf("expected unread count to be 0, got %d", refreshed.UnreadCount)
	}
}
