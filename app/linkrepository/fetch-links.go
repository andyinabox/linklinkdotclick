package linkrepository

import (
	"github.com/andyinabox/linkydink/app"
)

func (r *Repository) FetchLinks(userId uint) ([]app.Link, error) {
	var links []app.Link
	tx := r.withUserId(userId).Order("last_clicked").Find(&links)
	return links, tx.Error
}
