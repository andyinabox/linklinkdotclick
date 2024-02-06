package userservice

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
)

type Service struct {
	defaultUserId uint
	ur            app.UserRepository
	ts            app.TokenStore
	conf          *Config
}

type Config struct {
	DefaultUserEmail string
}

func New(ur app.UserRepository, ts app.TokenStore, conf *Config) *Service {
	if conf.DefaultUserEmail == "" {
		panic("no default user email provided")
	}

	// create default user
	user, err := ur.FetchUserByEmail(conf.DefaultUserEmail)
	if errors.Is(err, app.ErrNotFound) {
		user, err = ur.CreateUser(app.User{
			Email: conf.DefaultUserEmail,
		})
	}
	if err != nil {
		panic(err)
	}

	return &Service{user.ID, ur, ts, conf}
}
