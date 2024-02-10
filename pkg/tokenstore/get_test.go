package tokenstore

import (
	"errors"
	"testing"
	"time"

	"github.com/andyinabox/linkydink/test"
)

func Test_Get(t *testing.T) {
	db := test.NewInMemoryDb(t)
	store := New(db, &Config{})
	expectedId := uint(1)
	hash, err := store.Create(expectedId)
	if err != nil {
		t.Fatal(err.Error())
	}

	id, err := store.Get(hash)
	if err != nil {
		t.Fatal(err.Error())
	}

	if id != expectedId {
		t.Fatalf("expected id of %d, got %d", expectedId, id)
	}
}

func Test_GetExpiration(t *testing.T) {
	db := test.NewInMemoryDb(t)
	store := New(db, &Config{
		ExpireseIn: 1 * time.Millisecond,
	})
	expectedId := uint(1)
	hash, err := store.Create(expectedId)
	if err != nil {
		t.Fatal(err.Error())
	}

	time.Sleep(5 * time.Millisecond)

	_, err = store.Get(hash)
	if err == nil {
		t.Fatal("expected error, got none")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
