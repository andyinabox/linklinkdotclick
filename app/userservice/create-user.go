package userservice

import (
	"net/mail"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) CreateUser(email string) (*app.User, error) {
	// email validation
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, err
	}

	return s.r.CreateUser(app.User{
		Email: email,
	})
}