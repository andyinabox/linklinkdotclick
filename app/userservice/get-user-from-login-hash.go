package userservice

import "github.com/andyinabox/linkydink/app"

func (s *Service) GetUserFromLoginHash(hash string) (*app.User, error) {
	id, err := s.tokenStore.Get(hash)
	if err != nil {
		return nil, err
	}
	user, err := s.FetchUser(id)
	if err != nil {
		return nil, err
	}
	// if we are successfully using the hash, delete it
	s.tokenStore.Delete(hash)
	return user, err
}
