package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) RegisterLinkClick(userId uint, id uint, t time.Time) (link *app.Link, err error) {
	link, err = s.lr.FetchLink(userId, id)
	if err != nil {
		return
	}
	link.LastClicked = t
	return s.RefreshAndUpdateLink(userId, *link, true)
}
