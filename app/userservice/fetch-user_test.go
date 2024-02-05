package userservice

import (
	"testing"
)

func Test_FetchUser(t *testing.T) {
	s := NewUserService(t, &Config{
		// TODO: I don't think this should be necessary
		UserDbPath: "db/usr",
	})

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
