package userservice

import (
	"errors"
	"fmt"
	"path"
	"sync"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/feedhelper"
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

	lock := sync.RWMutex{}

	// try and get service from cache
	lock.RLock()
	linkService, ok := cachedLinkServices[user.ID]
	lock.RUnlock()

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
	feedHelper := feedhelper.New()
	linkService = linkservice.New(linkRepository, feedHelper)

	// cache result and return
	lock.Lock()
	cachedLinkServices[user.ID] = linkService
	lock.Unlock()

	return linkService, nil

}
