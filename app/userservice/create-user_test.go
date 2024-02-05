package userservice

import (
	"testing"
<<<<<<< HEAD
)

func Test_CreateUser(t *testing.T) {
	s := NewUserService(t, &Config{})
	invalidEmail := "example.com"
	_, err := s.CreateUser(invalidEmail)
=======

	"github.com/andyinabox/linkydink/app/userrepository"
)

func Test_CreateUser(t *testing.T) {
	r, err := userrepository.New(&userrepository.Config{
		DbFile: ":memory:",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s := New(&Config{}, r)

	invalidEmail := "example.com"
	_, err = s.CreateUser(invalidEmail)
>>>>>>> main
	if err == nil {
		t.Errorf("Expected email validation error for '%s', got none", invalidEmail)
	}

	validEmail := "test@example.com"
	_, err = s.CreateUser(validEmail)
	if err != nil {
		t.Fatal(err.Error())
	}

}
