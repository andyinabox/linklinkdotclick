package linkrepository

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

func New(conf *Config) *Repository {

	db, err := gorm.Open(sqlite.Open(conf.DbFile), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&app.Link{})
	if err != nil {
		panic(err)
	}
	return &Repository{db}
}
