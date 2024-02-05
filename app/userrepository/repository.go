package userrepository

import (
<<<<<<< HEAD
	"github.com/andyinabox/linkydink/app"
=======
	"io/fs"
	"os"
	"path"

	"github.com/andyinabox/linkydink/app"
	"github.com/glebarez/sqlite"
>>>>>>> main
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

<<<<<<< HEAD
func New(db *gorm.DB) *Repository {

	err := db.AutoMigrate(&app.User{})
	if err != nil {
		panic(err)
	}
	return &Repository{db}
=======
type Config struct {
	DbFile string
}

func New(conf *Config) (*Repository, error) {

	err := os.MkdirAll(path.Dir(conf.DbFile), fs.ModePerm)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(conf.DbFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&app.User{})
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
>>>>>>> main
}
