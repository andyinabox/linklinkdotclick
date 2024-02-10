package servicecontainer

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/mailservice"
)

type Container struct {
	userService app.UserService
	linkService app.LinkService
	logService  app.LogService
	mailService *mailservice.Service
}

func New(
	userService app.UserService,
	linkService app.LinkService,
	logService app.LogService,
	mailService *mailservice.Service,
) *Container {
	return &Container{
		userService,
		linkService,
		logService,
		mailService,
	}
}

func (c *Container) UserService() app.UserService {
	return c.userService
}

func (c *Container) LinkService() app.LinkService {
	return c.linkService
}

func (c *Container) LogService() app.LogService {
	return c.logService
}

func (c *Container) MailService() *mailservice.Service {
	return c.mailService
}
