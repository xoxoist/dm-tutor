package database

import "log"

type (
	Log interface {
		Printf(format string, v ...interface{})
		Verbose() bool
	}

	// logImpl implements the golang-migrate Logger interface
	logImpl struct {
		logger *log.Logger
	}
)

func NewLog(logger *log.Logger) Log {
	return &logImpl{logger: logger}
}

func (l *logImpl) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *logImpl) Verbose() bool {
	return true // or false, depending on whether you want verbose logging
}
