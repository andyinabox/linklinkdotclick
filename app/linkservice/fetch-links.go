package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

func (s *Service) FetchLinks(userId uint) ([]app.Link, error) {
	return s.lr.FetchLinks(userId)
}
