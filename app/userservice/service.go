package userservice

import (
	"github.com/andyinabox/linkydink/app"
)

const (
	defaultUserEmail  = "linkydink@linkydink.tld"
	defaultUserDbPath = ":memory:"
)

type Service struct {
	c          *Config
	r          app.UserRepository
	tokenStore app.TokenStore
}

type Config struct {
	UserDbPath       string
	DefaultUserEmail string
}

func New(r app.UserRepository, tokenStore app.TokenStore, c *Config) *Service {

	if c.DefaultUserEmail == "" {
		c.DefaultUserEmail = defaultUserEmail
	}
	if c.UserDbPath == "" {
		c.UserDbPath = defaultUserDbPath
	}
	return &Service{c, r, tokenStore}
}
