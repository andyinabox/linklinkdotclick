package userservice

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) EnsureDefaultUser() (*app.User, error) {
	user, err := s.r.FetchUserByEmail(s.c.DefaultUserEmail)
	if errors.Is(err, app.ErrNotFound) {
		user, err = s.r.CreateUser(app.User{
			Email: s.c.DefaultUserEmail,
		})
	}
	return user, err
}
