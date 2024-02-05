package userrepository

import (
	"github.com/andyinabox/linkydink/app"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {

	err := db.AutoMigrate(&app.User{})
	if err != nil {
		panic(err)
	}
	return &Repository{db}
}
