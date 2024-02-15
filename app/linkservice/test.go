package linkservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/linkrepository"
	"github.com/andyinabox/linkydink/pkg/logservice"
	"github.com/andyinabox/linkydink/test"
)

func NewLinkService(t *testing.T) app.LinkService {
	db := test.NewInMemoryDb(t)
	lr := linkrepository.New(db)
	log := logservice.New()
	return New(lr, log, &Config{})
}
