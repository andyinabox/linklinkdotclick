package app

type ServiceContainer interface {
	UserService() UserService
	LinkService() LinkService
	LogService() LogService
}
