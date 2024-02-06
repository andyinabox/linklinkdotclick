package userservice

import "github.com/andyinabox/linkydink/app"

func (s *Service) FetchUser(id uint) (*app.User, error) {
	return s.ur.FetchUser(id)
}
