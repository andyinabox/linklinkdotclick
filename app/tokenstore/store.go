package tokenstore

import (
	"time"

	"gorm.io/gorm"
)

var defaultExpiresIn = 10 * time.Minute

type Config struct {
	ExpireseIn time.Duration
}

type Store struct {
	db   *gorm.DB
	conf *Config
}

func New(db *gorm.DB, conf *Config) *Store {

	if conf.ExpireseIn == 0 {
		conf.ExpireseIn = defaultExpiresIn
	}

	err := db.AutoMigrate(&Token{})
	if err != nil {
		panic(err)
	}

	return &Store{db, conf}
}
