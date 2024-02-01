package linkrepository

import (
	"github.com/andyinabox/linkydink/app"
)

func (r *Repository) CreateLink(link app.Link) (*app.Link, error) {
	tx := r.db.Create(&link)
	return &link, tx.Error
}
