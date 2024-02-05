package userservice

import (
	"testing"
<<<<<<< HEAD
)

func Test_FetchUser(t *testing.T) {
	s := NewUserService(t, &Config{
		// TODO: I don't think this should be necessary
		UserDbPath: "db/usr",
	})
=======

	"github.com/andyinabox/linkydink/app/userrepository"
)

func Test_FetchUser(t *testing.T) {
	r, err := userrepository.New(&userrepository.Config{
		DbFile: ":memory:",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s := New(&Config{
		UserDbPath: "db/usr",
	}, r)
>>>>>>> main

	validEmail := "test@example.com"
	createdUser, err := s.CreateUser(validEmail)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = s.FetchUser(createdUser.ID)
	if err != nil {
		t.Fatal(err.Error())
	}

}
