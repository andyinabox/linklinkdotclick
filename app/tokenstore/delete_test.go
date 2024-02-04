package tokenstore

import (
	"errors"
	"testing"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/test"
)

func Test_Delete(t *testing.T) {
	db := test.NewInMemoryDb(t)
	store := New(db, &Config{})
	hash, err := store.Create(1)
	if err != nil {
		t.Fatal(err.Error())
	}

	err = store.Delete(hash)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = store.Get(hash)
	if err == nil {
		t.Fatal("expected error, got none")
	}
	if !errors.Is(err, app.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
