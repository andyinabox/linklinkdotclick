package linkservice

import (
	"time"

	"github.com/andyinabox/linkydink/app"
)

type Config struct {
	LinkRefreshBuffer time.Duration
}

type Service struct {
	lr   app.LinkRepository
	log  app.LogService
	conf *Config
}

func New(lr app.LinkRepository, log app.LogService, conf *Config) *Service {
	return &Service{lr, log, conf}
}
