package userrepository

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
	"gorm.io/gorm"
)

func (r *Repository) FetchUserByEmail(email string) (*app.User, error) {
	var user app.User
	var err error
	tx := r.db.Where("email = ?", email).First(&user)
	err = tx.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = app.ErrNotFound
	}

	return &user, err
}
