package userservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/userrepository"
)

func Test_EnsureDefaultUser(t *testing.T) {
	r, err := userrepository.New(&userrepository.Config{
		DbFile: ":memory:",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	email := "default@user.com"
	s := New(&Config{
		DefaultUserEmail: email,
	}, r)

	user, err := s.EnsureDefaultUser()
	if err != nil {
		t.Fatal(err.Error())
	}
	if user == nil {
		t.Fatal("expected default user, recieved nil")
	}
	if user.Email != email {
		t.Errorf("expected default user to have email %s, got %s", email, user.Email)
	}
}