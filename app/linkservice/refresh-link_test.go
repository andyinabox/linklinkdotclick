package linkservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/linkrepository"
	"github.com/andyinabox/linkydink/test/fixtures"
)

func Test_RefreshLink(t *testing.T) {
	lr, err := linkrepository.New(&linkrepository.Config{
		DbFile: ":memory:",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	ls := New(lr)

	link := fixtures.LinkJustClicked()
	refreshed, err := ls.refreshLink(link)
	if err != nil {
		t.Fatal(err.Error())
	}
	if refreshed.UnreadCount != 0 {
		t.Errorf("expected unread count to be 0, got %d", refreshed.UnreadCount)
	}
}
