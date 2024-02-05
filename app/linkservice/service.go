package linkservice

import (
	"github.com/andyinabox/linkydink/app"
	"github.com/andyinabox/linkydink/pkg/feedfinder"
)

type Service struct {
	lr app.LinkRepository
	fr *feedfinder.Reader
}

func New(lr app.LinkRepository) *Service {
	fr := feedfinder.New()
	return &Service{lr, fr}
}
