package userservice

import "github.com/andyinabox/linkydink/app"

func (s *Service) FetchDefaultUser() (*app.User, error) {
	id := s.GetDefaultUserId()
	return s.FetchUser(id)
}
