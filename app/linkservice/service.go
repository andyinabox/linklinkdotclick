package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

type Service struct {
	lr app.LinkRepository
	fh app.FeedHelper
}

func New(lr app.LinkRepository, fh app.FeedHelper) *Service {
	return &Service{lr, fh}
}
