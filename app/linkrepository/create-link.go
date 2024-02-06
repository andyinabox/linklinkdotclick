package linkrepository

import (
	"github.com/andyinabox/linkydink/app"
)

func (r *Repository) CreateLink(link app.Link) (*app.Link, error) {
	tx := r.withUserId(link.UserID).Create(&link)
	return &link, tx.Error
}
