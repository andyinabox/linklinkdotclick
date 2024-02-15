package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

func (s *Service) UpdateLink(userId uint, link app.Link, refresh bool) (*app.Link, error) {
	if refresh {
		return s.RefreshAndUpdateLink(userId, link)
	}
	return s.lr.UpdateLink(link)
}
