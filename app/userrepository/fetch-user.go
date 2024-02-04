package userrepository

import "github.com/andyinabox/linkydink/app"

func (r *Repository) FetchUser(id uint) (*app.User, error) {
	var user app.User
	tx := r.db.First(&user, id)
	return &user, tx.Error
}
