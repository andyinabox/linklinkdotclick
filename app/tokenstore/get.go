package tokenstore

import (
	"errors"
	"time"

	"github.com/andyinabox/linkydink/app"
	"gorm.io/gorm"
)

func (s *Store) Get(hash string) (userId uint, err error) {
	var token Token

	tx := s.db.Where("hash = ? AND expires_at > ?", hash, time.Now()).First(&token)
	err = tx.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = app.ErrNotFound
		}
		return
	}

	userId = token.UserID

	return
}
