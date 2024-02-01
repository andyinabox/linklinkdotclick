package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

func (s *Service) FetchLink(id uint, refresh bool) (*app.Link, error) {

	link, err := s.lr.FetchLink(id)
	if err != nil {
		return nil, err
	}

	if refresh {
		// refresh link feed data
		link, err = s.refreshLink(*link)
		if err != nil {
			return nil, err
		}
		// save update link to db
		link, err = s.lr.UpdateLink(*link)
		if err != nil {
			return nil, err
		}
	}

	return link, nil
}
