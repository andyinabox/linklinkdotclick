package userservice

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) UpdateUser(id uint, user app.User) (*app.User, error) {
	if id != user.ID {
		return nil, errors.New("unmatching ids in update request")
	}

	// disallow upsert
	_, err := s.ur.FetchUser(id)
	if err != nil {
		return nil, err
	}

	return s.ur.UpsertUser(user)
}
