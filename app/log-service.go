package app

import "log"

type LogService interface {
	Info() *log.Logger
	Warn() *log.Logger
	Error() *log.Logger
}
