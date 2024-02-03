package userservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/userrepository"
)

func Test_CreateUser(t *testing.T) {
	r, err := userrepository.New(&userrepository.Config{
		DbFile: ":memory:",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	s := New(&Config{
		UserDbPath: "db/usr",
	}, r)

	invalidEmail := "example.com"
	_, err = s.CreateUser(invalidEmail)
	if err == nil {
		t.Errorf("Expected email validation error for '%s', got none", invalidEmail)
	}

	validEmail := "test@example.com"
	_, err = s.CreateUser(validEmail)
	if err != nil {
		t.Fatal(err.Error())
	}

}
