package userrepository

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type Config struct {
	DbFile string
}

func New(conf *Config) (*Repository, error) {
	db, err := gorm.Open(sqlite.Open(conf.DbFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&app.User{})
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}
