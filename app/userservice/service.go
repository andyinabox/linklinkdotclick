package userservice

import (
	"github.com/andyinabox/linkydink/app"
)

type Service struct {
	ur   app.UserRepository
	ts   app.TokenStore
	conf *Config
}

type Config struct {
	DefaultUserEmail string
}

func New(ur app.UserRepository, ts app.TokenStore, conf *Config) *Service {
	if conf.DefaultUserEmail == "" {
		panic("no default user email provided")
	}
	return &Service{ur, ts, conf}
}
