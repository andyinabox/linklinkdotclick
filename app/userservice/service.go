package userservice

import "github.com/andyinabox/linkydink/app"

type Service struct {
	c *Config
	r app.UserRepository
}

type Config struct {
	UserDbPath string
}

func New(c *Config, r app.UserRepository) *Service {
	return &Service{c, r}
}
