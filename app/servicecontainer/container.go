package servicecontainer

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/mailservice"
)

type Container struct {
	us app.UserService
	ls app.LinkService
	ms *mailservice.Service
}

func New(us app.UserService, ls app.LinkService, ms *mailservice.Service) *Container {
	return &Container{us, ls, ms}
}

func (c *Container) UserService() app.UserService {
	return c.us
}

func (c *Container) LinkService() app.LinkService {
	return c.ls
}

func (c *Container) MailService() *mailservice.Service {
	return c.ms
}
