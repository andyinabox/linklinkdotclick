package linkservice

import "github.com/andyinabox/linkydink/app"

func (s *Service) DeleteLink(userId uint, id uint) (uint, error) {
	if userId == 0 {
		return 0, app.ErrMissingUserId
	}
	return s.lr.DeleteLink(id)
}
