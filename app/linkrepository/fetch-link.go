package linkrepository

import (
	"github.com/andyinabox/linkydink/app"
)

func (r *Repository) FetchLink(id uint) (*app.Link, error) {
	var link app.Link
	tx := r.db.First(&link, id)
	return &link, tx.Error
}
