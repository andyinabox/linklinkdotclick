package userrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
)

func Test_FetchUser(t *testing.T) {
	r, err := New(&Config{":memory:"})
	if err != nil {
		t.Fatal(err.Error())
	}

	// test fetching non-existent record
	fetchedUser, err := r.FetchUser(99)
	if err == nil {
		t.Errorf("expected fetching non-existent record to fail, got %v", fetchedUser)
	}

	link := app.User{
		Email: "test@example.com",
	}
	createdUser, err := r.CreateUser(link)
	if err != nil {
		t.Fatal(err.Error())
	}
	fetchedUser, err = r.FetchUser(createdUser.ID)
	if err != nil {
		t.Fatal(err.Error())
	}
	if fetchedUser.ID != createdUser.ID {
		t.Errorf("expected fetched link ID to be equal to original link: %v, %v", fetchedUser.ID, createdUser.ID)
	}
}
