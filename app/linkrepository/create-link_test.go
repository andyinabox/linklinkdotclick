package linkrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/test/fixtures"
)

func Test_CreateLink(t *testing.T) {
	r := NewLinkRepository(t)

	{
		link := fixtures.LinkBeforeInsert()
		createdLink, err := r.CreateLink(link)
		if err != nil {
			t.Fatal(err.Error())
		}
		if createdLink.ID == 0 {
			t.Fatal("expected id to be non-zero value")
		}
	}

	{
		link := fixtures.LinkBeforeInsert()
		link.UserID = 0
		_, err := r.CreateLink(link)
		if err == nil {
			t.Fatal("expected error createing link with zero-value UserID")
		}
	}

}
