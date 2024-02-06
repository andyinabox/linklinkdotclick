package userservice

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) EnsureDefaultUser() (*app.User, error) {
	user, err := s.ur.FetchUserByEmail(s.conf.DefaultUserEmail)
	if errors.Is(err, app.ErrNotFound) {
		user, err = s.ur.CreateUser(app.User{
			Email: s.conf.DefaultUserEmail,
		})
	}
	return user, err
}
