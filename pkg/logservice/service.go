package logservice

import (
	"log"
	"os"
)

type Logger struct {
	errlogger  *log.Logger
	warnlogger *log.Logger
	infologger *log.Logger
}

func New() *Logger {
	errlogger := log.New(os.Stdout, "ERROR: ", log.Llongfile)
	warnlogger := log.New(os.Stdout, "WARN: ", log.Llongfile)
	infologger := log.New(os.Stdout, "INFO: ", log.Llongfile)
	return &Logger{
		errlogger,
		warnlogger,
		infologger,
	}
}

func (l *Logger) Error() *log.Logger {
	return l.errlogger
}

func (l *Logger) Warn() *log.Logger {
	return l.warnlogger
}

func (l *Logger) Info() *log.Logger {
	return l.infologger
}
