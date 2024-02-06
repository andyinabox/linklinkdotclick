package linkrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/test/fixtures"
)

func Test_FetchLink(t *testing.T) {
	r := NewLinkRepository(t)

	{ // test fetching non-existent record
		fetchedLink, err := r.FetchLink(1, 99)
		if err == nil {
			t.Errorf("expected fetching non-existent record to fail, got %v", fetchedLink)
		}

		link := fixtures.LinkBeforeInsert()
		createdLink, err := r.CreateLink(link)
		if err != nil {
			t.Fatal(err.Error())
		}
		fetchedLink, err = r.FetchLink(1, createdLink.ID)
		if err != nil {
			t.Fatal(err.Error())
		}
		if fetchedLink.ID != createdLink.ID {
			t.Errorf("expected fetched link ID to be equal to original link: %v, %v", fetchedLink.ID, link.ID)
		}
	}
	{
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic with missing user ID")
			}
		}()
		_, _ = r.FetchLink(0, 1)
	}
}
