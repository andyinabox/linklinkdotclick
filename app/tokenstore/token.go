package tokenstore

import "time"

type Token struct {
	ID        uint
	Hash      string `gorm:"uniqueIndex"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time `gorm:"index"`
}
