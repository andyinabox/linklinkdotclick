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
