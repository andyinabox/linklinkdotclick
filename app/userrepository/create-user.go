package userrepository

import "github.com/andyinabox/linkydink/app"

func (r *Repository) CreateUser(user app.User) (*app.User, error) {
	tx := r.db.Create(&user)
	return &user, tx.Error
}
