package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

type Service struct {
	lr app.LinkRepository
	fs app.FeedService
}

func New(lr app.LinkRepository, fs app.FeedService) *Service {
	return &Service{lr, fs}
}
