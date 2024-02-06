package userservice

import (
	"testing"
)

func Test_FetchUser(t *testing.T) {
	s := NewUserService(t, &Config{})

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
