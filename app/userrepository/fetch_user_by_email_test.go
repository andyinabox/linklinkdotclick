package userrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
<<<<<<< HEAD
	"github.com/andyinabox/linkydink/test"
)

func Test_FetchUserByEmail(t *testing.T) {
	db := test.NewInMemoryDb(t)
	r := New(db)
=======
)

func Test_FetchUserByEmail(t *testing.T) {
	r, err := New(&Config{":memory:"})
	if err != nil {
		t.Fatal(err.Error())
	}
>>>>>>> main

	email := "test@example.com"

	// test fetching non-existent email
	fetchedUser, err := r.FetchUserByEmail(email)
	if err == nil {
		t.Errorf("expected fetching non-existent record to fail, got %v", fetchedUser)
	}

<<<<<<< HEAD
	user := app.User{
		Email: email,
	}
	_, err = r.CreateUser(user)
=======
	link := app.User{
		Email: email,
	}
	_, err = r.CreateUser(link)
>>>>>>> main
	if err != nil {
		t.Fatal(err.Error())
	}
	_, err = r.FetchUserByEmail(email)
	if err != nil {
		t.Fatal(err.Error())
	}
}
