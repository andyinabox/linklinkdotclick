package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

func (s *Service) FetchLinks(userId uint) ([]app.Link, error) {
	links, err := s.lr.FetchLinks(userId)
	if err != nil {
		return nil, err
	}

	for _, link := range links {
		go s.refreshLink(userId, link)
	}

	return s.lr.FetchLinks(userId)
}

func (s *Service) refreshLink(userId uint, link app.Link) {

	refreshedLink, err := s.RefreshLink(userId, link)
	if err != nil {
		s.log.Error().Println(err.Error())
	}
	// save update link to db
	updatedLink, err := s.lr.UpdateLink(*refreshedLink)
	if err != nil {
		s.log.Error().Println(err.Error())
	}

	s.log.Info().Printf("Refreshed link %s in goroutine", updatedLink.SiteName)
}
