package linkservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/feedservice"
	"github.com/andyinabox/linkydink/app/linkrepository"
	"github.com/andyinabox/linkydink/test"
)

func NewLinkService(t *testing.T) app.LinkService {
	db := test.NewInMemoryDb(t)
	lr := linkrepository.New(db)
	fs := feedservice.New()
	return New(lr, fs)
}
