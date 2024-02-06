package userservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/tokenstore"
	"github.com/andyinabox/linkydink/app/userrepository"
	"github.com/andyinabox/linkydink/test"
)

func NewUserService(t *testing.T, conf *Config) app.UserService {
	if conf.DefaultUserEmail == "" {
		conf.DefaultUserEmail = "test2@example.com"
	}
	db := test.NewInMemoryDb(t)
	r := userrepository.New(db)
	ts := tokenstore.New(db, &tokenstore.Config{})
	return New(r, ts, conf)
}
