package linkrepository

import (
	"io/fs"
	"os"
	"path"

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

	err := os.MkdirAll(path.Dir(conf.DbFile), fs.ModePerm)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(conf.DbFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&app.Link{})
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}
