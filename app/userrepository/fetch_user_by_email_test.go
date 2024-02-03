package userrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
)

func Test_FetchUserByEmail(t *testing.T) {
	r, err := New(&Config{":memory:"})
	if err != nil {
		t.Fatal(err.Error())
	}

	email := "test@example.com"

	// test fetching non-existent email
	fetchedUser, err := r.FetchUserByEmail(email)
	if err == nil {
		t.Errorf("expected fetching non-existent record to fail, got %v", fetchedUser)
	}

	link := app.User{
		Email: email,
	}
	_, err = r.CreateUser(link)
	if err != nil {
		t.Fatal(err.Error())
	}
	_, err = r.FetchUserByEmail(email)
	if err != nil {
		t.Fatal(err.Error())
	}
}
