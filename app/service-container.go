package app

import (
	"github.com/andyinabox/linkydink/pkg/mailservice"
)

type ServiceContainer interface {
	UserService() UserService
	LinkService() LinkService
	MailService() *mailservice.Service
	LogService() LogService
}
