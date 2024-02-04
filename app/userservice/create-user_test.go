package userservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/tokenstore"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/test"
)

func Test_CreateUser(t *testing.T) {
	db := test.NewInMemoryDb(t)
	r := userrepository.New(db)
	s := New(&Config{}, r, tokenstore.New(db, &tokenstore.Config{}))

	invalidEmail := "example.com"
	_, err := s.CreateUser(invalidEmail)
	if err == nil {
		t.Errorf("Expected email validation error for '%s', got none", invalidEmail)
	}

	validEmail := "test@example.com"
	_, err = s.CreateUser(validEmail)
	if err != nil {
		t.Fatal(err.Error())
	}

}
