package linkrepository

import (
	"github.com/andyinabox/linkydink/app"
)

func (r *Repository) FetchLinks() ([]app.Link, error) {
	var links []app.Link
	tx := r.db.Order("last_clicked").Find(&links)
	return links, tx.Error
}
