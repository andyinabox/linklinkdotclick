package linkrepository

import (
	"testing"
	"time"

	"github.com/andyinabox/linkydink/test/fixtures"
)

func Test_UpdateLink(t *testing.T) {
	r := NewLinkRepository(t)

	{
		link := fixtures.LinkBeforeInsert()
		createdLink, err := r.CreateLink(link)
		if err != nil {
			t.Fatal(err.Error())
		}

		updateLink := *createdLink
		updateLink.LastClicked = time.Now()
		updateLink.UnreadCount = 0

		result, err := r.UpdateLink(updateLink)
		if err != nil {
			t.Fatal(err.Error())
		}
		if result.ID != createdLink.ID {
			t.Errorf("expected result link ID to be equal to original link: %v, %v", result.ID, link.ID)
		}
		if result.LastClicked == createdLink.LastClicked {
			t.Errorf("expected result LastClicked value to be different: %v, %v", result.LastClicked, link.LastClicked)
		}
	}

	{
		link := fixtures.LinkBeforeInsert()
		link.UserID = 0
		_, err := r.UpdateLink(link)
		if err == nil {
			t.Fatal("expected error createing link with zero-value UserID")
		}
	}

}
