package userservice

<<<<<<< HEAD
import (
	"github.com/andyinabox/linkydink/app"
)
=======
import "github.com/andyinabox/linkydink/app"
>>>>>>> main

const (
	defaultUserEmail  = "linkydink@linkydink.tld"
	defaultUserDbPath = ":memory:"
)

type Service struct {
<<<<<<< HEAD
	c          *Config
	r          app.UserRepository
	tokenStore app.TokenStore
=======
	c *Config
	r app.UserRepository
>>>>>>> main
}

type Config struct {
	UserDbPath       string
	DefaultUserEmail string
}

<<<<<<< HEAD
func New(r app.UserRepository, tokenStore app.TokenStore, c *Config) *Service {
=======
func New(c *Config, r app.UserRepository) *Service {
>>>>>>> main

	if c.DefaultUserEmail == "" {
		c.DefaultUserEmail = defaultUserEmail
	}
	if c.UserDbPath == "" {
		c.UserDbPath = defaultUserDbPath
	}
<<<<<<< HEAD
	return &Service{c, r, tokenStore}
=======
	return &Service{c, r}
>>>>>>> main
}
