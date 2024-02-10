package servicecontainer

import (
	"github.com/andyinabox/linkydink/app"
)

type Container struct {
	userService app.UserService
	linkService app.LinkService
	logService  app.LogService
}

func New(
	userService app.UserService,
	linkService app.LinkService,
	logService app.LogService,
) *Container {
	return &Container{
		userService,
		linkService,
		logService,
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
