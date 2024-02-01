package linkrepository

import (
	"testing"
	"time"

	"github.com/andyinabox/linkydink/test/fixtures"
)

func Test_UpdateLink(t *testing.T) {
	r := New(&Config{":memory:"})

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

// func Test_UpdateLinkInsert(t *testing.T) {
// 	r := New(&Config{":memory:"})

// 	link := test.LinkBeforeInsert()
// 	link.ID = 1
// 	_, err := r.UpdateLink(*link)
// 	if err == nil {
// 		t.Fatal("expeted error when attempting to update non-existing record")
// 	}
// 	if errors.Is(err, app.ErrNotFound) {
// 		t.Fatalf("expected ErrNotFound error, got %v", err)
// 	}
// }
