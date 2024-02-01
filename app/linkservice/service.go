package linkservice

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/app/feedreader"
)

type Service struct {
	lr app.LinkRepository
	fr *feedreader.Reader
}

func New(lr app.LinkRepository) *Service {
	fr := feedreader.New()
	return &Service{lr, fr}
}
