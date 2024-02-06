package linkrepository

import (
	"errors"

	"github.com/andyinabox/linkydink/app"
	"gorm.io/gorm"
)

func (r *Repository) FetchLink(userId uint, id uint) (*app.Link, error) {
	var link app.Link
	tx := r.withUserId(userId).First(&link, id)
	err := tx.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = app.ErrNotFound
	}
	return &link, err
}
