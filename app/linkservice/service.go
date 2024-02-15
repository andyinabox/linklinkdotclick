package linkservice

import (
	"github.com/andyinabox/linkydink/app"
)

type Service struct {
	lr  app.LinkRepository
	log app.LogService
}

func New(lr app.LinkRepository, log app.LogService) *Service {
	return &Service{lr, log}
}
