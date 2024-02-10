package tokenstore

import (
	"testing"

	"github.com/andyinabox/linkydink/test"
)

func Test_Create(t *testing.T) {
	db := test.NewInMemoryDb(t)
	store := New(db, &Config{})
	_, err := store.Create(1)
	if err != nil {
		t.Fatal(err.Error())
	}
}
