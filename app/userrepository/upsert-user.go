package userrepository

import "github.com/andyinabox/linkydink/app"

func (r *Repository) UpsertUser(user app.User) (*app.User, error) {
	tx := r.db.Save(&user)
	return &user, tx.Error
}
