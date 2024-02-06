package userservice

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) FetchOrCreateUserByEmail(email string) (user *app.User, err error) {
	user, err = s.ur.FetchUserByEmail(email)
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			user, err = s.ur.CreateUser(app.User{Email: email})
		}
	}
	return
}
