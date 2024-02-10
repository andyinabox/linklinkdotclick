package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

type Service struct {
	lr  app.LinkRepository
	fs  app.FeedService
	log app.LogService
}

func New(lr app.LinkRepository, fs app.FeedService, log app.LogService) *Service {
	return &Service{lr, fs, log}
}
