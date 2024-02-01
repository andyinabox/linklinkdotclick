package linkrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/test"
)

func Test_FetchLink(t *testing.T) {
	r := New(&Config{":memory:"})

	// test fetching non-existent record
	fetchedLink, err := r.FetchLink(99)
	if err == nil {
		t.Errorf("expected fetching non-existent record to fail, got %v", fetchedLink)
	}

	link := test.LinkBeforeInsert()
	link, err = r.CreateLink(*link)
	if err != nil {
		t.Fatal(err.Error())
	}
	fetchedLink, err = r.FetchLink(link.ID)
	if err != nil {
		t.Fatal(err.Error())
	}
	if fetchedLink.ID != link.ID {
		t.Errorf("expected fetched link ID to be equal to original link: %v, %v", fetchedLink.ID, link.ID)
	}
}
