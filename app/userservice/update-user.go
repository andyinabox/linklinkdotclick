package userservice

import (
	"github.com/andyinabox/linkydink/app"
)

func (s *Service) UpdateUser(id uint, user app.User) (*app.User, error) {

	user.ID = id

	// disallow upsert
	fetchedUser, err := s.ur.FetchUser(id)
	if err != nil {
		return nil, err
	}

	user.Email = fetchedUser.Email

	return s.ur.UpsertUser(user)
}
