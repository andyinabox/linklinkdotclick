package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

func (s *Service) FetchLinks() ([]app.Link, error) {
	return s.lr.FetchLinks()
}
