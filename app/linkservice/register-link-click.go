package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
)

func (s *Service) RegisterLinkClick(userId uint, id uint) (link *app.Link, err error) {
	link, err = s.lr.FetchLink(userId, id)
	if err != nil {
		return
	}
	link.LastClicked = time.Now()
	return s.RefreshAndUpdateLink(userId, *link)
}
