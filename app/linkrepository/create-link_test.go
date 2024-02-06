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
		defer func() {
			if r := recover(); r == nil {
				t.Error("creating link without userId did not panic")
			}
		}()
		_, _ = r.CreateLink(fixtures.LinkBeforeInsertNoUserID())
	}

}
