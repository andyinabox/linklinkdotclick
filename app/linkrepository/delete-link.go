package linkrepository

import "github.com/andyinabox/linkydink/app"

func (r *Repository) DeleteLink(id uint) (uint, error) {
	tx := r.db.Delete(&app.Link{}, id)
	return id, tx.Error
}
