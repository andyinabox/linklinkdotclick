package userservice

import (
	"testing"
)

func Test_EnsureDefaultUser(t *testing.T) {
	email := "default@user.com"

	s := NewUserService(t, &Config{
		DefaultUserEmail: email,
	})

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
