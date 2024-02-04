package userservice

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) FetchOrCreateUserByEmail(email string) (user *app.User, err error) {
	user, err = s.r.FetchUserByEmail(email)
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
			user, err = s.r.CreateUser(app.User{Email: email})
		}
	}
	return
}
