package userservice

import "github.com/andyinabox/linkydink/app"

const (
	defaultUserEmail  = "linkydink@linkydink.tld"
	defaultUserDbPath = "db/usr"
)

type Service struct {
	c *Config
	r app.UserRepository
}

type Config struct {
	UserDbPath       string
	DefaultUserEmail string
}

func New(c *Config, r app.UserRepository) *Service {

	if c.DefaultUserEmail == "" {
		c.DefaultUserEmail = defaultUserEmail
	}
	if c.UserDbPath == "" {
		c.UserDbPath = defaultUserDbPath
	}
	return &Service{c, r}
}
