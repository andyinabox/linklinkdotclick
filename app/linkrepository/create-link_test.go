package linkrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/test/fixtures"
)

func Test_CreateLink(t *testing.T) {
	r, err := New(&Config{":memory:"})
	if err != nil {
		t.Fatal(err.Error())
	}
	link := fixtures.LinkBeforeInsert()
	createdLink, err := r.CreateLink(link)
	if err != nil {
		t.Fatal(err.Error())
	}
	if createdLink.ID == 0 {
		t.Fatal("expected id to be non-zero value")
	}

}
