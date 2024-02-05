package linkservice

import (
	"testing"

	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/feedhelper"
	"github.com/andyinabox/linkydink/app/linkrepository"
)

func NewLinkService(t *testing.T) app.LinkService {
	lr, err := linkrepository.New(&linkrepository.Config{
		DbFile: ":memory:",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fh := feedhelper.New()
	return New(lr, fh)
}
