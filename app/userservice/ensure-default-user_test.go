package userservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/tokenstore"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/test"
)

func Test_EnsureDefaultUser(t *testing.T) {
	email := "default@user.com"

	db := test.NewInMemoryDb(t)
	r := userrepository.New(db)
	s := New(&Config{
		DefaultUserEmail: email,
	}, r, tokenstore.New(db, &tokenstore.Config{}))

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
