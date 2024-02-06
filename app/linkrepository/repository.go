package linkrepository

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type Config struct {
	DbFile string
}

func New(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) withUserId(userId uint) *gorm.DB {
	if userId == 0 {
		panic("zero-value user id")
	}
	return r.db.Where("user_id = ?", userId)
}
