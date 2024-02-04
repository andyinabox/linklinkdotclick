package userservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/tokenstore"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/test"
)

func Test_FetchUser(t *testing.T) {
	db := test.NewInMemoryDb(t)
	r := userrepository.New(db)
	s := New(&Config{
		UserDbPath: "db/usr",
	}, r, tokenstore.New(db, &tokenstore.Config{}))

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
