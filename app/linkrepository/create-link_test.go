package linkrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/test"
)

func Test_CreateLink(t *testing.T) {
	r := New(&Config{":memory:"})
	link := test.LinkBeforeInsert()
	link, err := r.CreateLink(*link)
	if err != nil {
		t.Fatal(err.Error())
	}
	if link.ID == 0 {
		t.Fatal("expected id to be non-zero value")
	}

}
