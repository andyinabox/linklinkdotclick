package userservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app/tokenstore"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/test"
)

func Test_New(t *testing.T) {
	db := test.NewInMemoryDb(t)
	r := userrepository.New(db)
	_ = New(&Config{
		UserDbPath: "db/usr",
	}, r, tokenstore.New(db, &tokenstore.Config{}))
}
