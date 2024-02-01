package linkservice

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) UpdateLink(id uint, link app.Link, refresh bool) (*app.Link, error) {
	if id != link.ID {
		return nil, errors.New("unmatching ids in update request")
	}
	updatedLink, err := s.lr.UpdateLink(link)
	if err != nil {
		return nil, err
	}

	if refresh {
		updatedLink, err = s.refreshLink(*updatedLink)
		if err != nil {
			return nil, err
		}
		updatedLink, err = s.lr.UpdateLink(*updatedLink)
		if err != nil {
			return nil, err
		}
	}

	return updatedLink, nil
}
