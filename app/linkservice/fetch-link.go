package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

func (s *Service) FetchLink(userId uint, id uint, refresh bool) (link *app.Link, err error) {

	if refresh {
		link, err = s.lr.FetchLink(userId, id)
		if err != nil {
			return nil, err
		}

		// refresh link feed data
		return s.RefreshAndUpdateLink(userId, *link, false)
	}

	return s.lr.FetchLink(userId, id)
}
