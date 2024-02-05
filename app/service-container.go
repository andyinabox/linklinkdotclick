package app

import (
	"github.com/andyinabox/linkydink/pkg/mailservice"
)

type ServiceContainer interface {
	UserService() UserService
	// this will get the link service for the default user
	DefaultLinkService() LinkService
	MailService() *mailservice.Service
}
