package userrepository

import (
	"testing"

	"github.com/andyinabox/linkydink/test"
)

func Test_New(t *testing.T) {
	db := test.NewInMemoryDb(t)
	_ = New(db)
}
