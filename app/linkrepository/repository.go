package linkrepository

import (
	"github.com/andyinabox/linkydink/app"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type Config struct {
	DbFile string
}

func New(db *gorm.DB) *Repository {

	err := db.AutoMigrate(&app.Link{})
	if err != nil {
		panic(err)
	}

	return &Repository{db}
}

func (r *Repository) withUserId(userId uint) *gorm.DB {
	if userId == 0 {
		panic("zero-value user id")
	}
	return r.db.Where("user_id = ?", userId)
}
