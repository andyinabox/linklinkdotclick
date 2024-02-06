package linkrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/test"
)

func NewLinkRepository(t *testing.T) *Repository {
	db := test.NewInMemoryDb(t)
	err := db.AutoMigrate(&app.Link{})
	if err != nil {
		panic(err)
	}
	return New(db)
}
