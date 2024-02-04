package userservice

import (
	"testing"

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
