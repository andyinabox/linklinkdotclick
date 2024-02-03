package userservice

import (
	"errors"
	"fmt"
	"path"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/linkrepository"
	"github.com/andyinabox/linkydink/app/linkservice"
)

var cachedLinkServices = map[uint]app.LinkService{}

func (s *Service) getUserDbFilePath(id uint) string {
	if s.c.UserDbPath == ":memory:" {
		return fmt.Sprintf("file:%d?mode=memory", id)
	}
	return path.Join(s.c.UserDbPath, fmt.Sprintf("%d.db", id))
}

func (s *Service) GetUserLinkService(user *app.User) (app.LinkService, error) {
	if user.ID == 0 {
		return nil, errors.New("invalid user id: 0")
	}

	// try and get service from cache
	linkService, ok := cachedLinkServices[user.ID]
	if ok {
		return linkService, nil
	}

	// create new repository and service
	linkRepository, err := linkrepository.New(&linkrepository.Config{
		DbFile: s.getUserDbFilePath(user.ID),
	})
	if err != nil {
		return nil, err
	}
	linkService = linkservice.New(linkRepository)

	// cache result and return
	cachedLinkServices[user.ID] = linkService
	return linkService, nil

}
