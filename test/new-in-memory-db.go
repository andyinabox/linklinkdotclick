package test

import (
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewInMemoryDb(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err.Error())
	}
	return db
}
