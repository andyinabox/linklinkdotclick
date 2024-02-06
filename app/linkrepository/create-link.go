package linkrepository

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
)

func (r *Repository) CreateLink(link app.Link) (*app.Link, error) {
	if link.UserID == 0 {
		return nil, errors.New("zero-value user id")
	}
	tx := r.db.Create(&link)
	return &link, tx.Error
}
