package tokenstore

import (
	"time"
)

func (s *Store) New(userId uint) (hash string, err error) {
	hash, err = generateHash(32)
	if err != nil {
		return
	}

	token := &Token{
		Hash:      hash,
		UserID:    userId,
		ExpiresAt: time.Now().Add(s.conf.ExpireseIn),
	}

	tx := s.db.Create(&token)
	err = tx.Error

	return
}
