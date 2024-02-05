package userservice

import "github.com/andyinabox/linkydink/app"

func (s *Service) GetLoginHashForUser(user *app.User) (string, error) {
	return s.tokenStore.Create(user.ID)
}
