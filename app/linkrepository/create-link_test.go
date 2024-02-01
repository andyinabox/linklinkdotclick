package linkrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/test/fixtures"
)

func Test_CreateLink(t *testing.T) {
	r := New(&Config{":memory:"})
	link := fixtures.LinkBeforeInsert()
	createdLink, err := r.CreateLink(link)
	if err != nil {
		t.Fatal(err.Error())
	}
	if createdLink.ID == 0 {
		t.Fatal("expected id to be non-zero value")
	}

}
