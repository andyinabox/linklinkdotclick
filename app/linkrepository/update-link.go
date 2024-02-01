package linkrepository

import (
	"github.com/andyinabox/linkydink/app"
)

func (r *Repository) UpdateLink(link app.Link) (*app.Link, error) {
	tx := r.db.Save(&link)
	return &link, tx.Error
}
