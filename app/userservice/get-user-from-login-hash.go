package userservice

import "github.com/andyinabox/linkydink/app"

func (s *Service) GetUserFromLoginHash(hash string) (*app.User, error) {
	id, err := s.ts.Get(hash)
	if err != nil {
		return nil, err
	}
	user, err := s.FetchUser(id)
	if err != nil {
		return nil, err
	}
	// if we are successfully using the hash, delete it
	s.ts.Delete(hash)
	return user, err
}
