package linkrepository

import "github.com/andyinabox/linkydink/app"

func (r *Repository) DeleteLink(userId uint, id uint) (uint, error) {
	tx := r.withUserId(userId).Delete(&app.Link{}, id)
	return id, tx.Error
}
