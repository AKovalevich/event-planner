package app

import (
	"github.com/Sirupsen/logrus"
)

// logger wraps logrus.Logger so that it can log messages sharing a common set of fields.
type logger struct {
	logger *logrus.Logger
	fields logrus.Fields
}

var log *logger

//
func Logger() *logger {
	once.Do(func() {
		log = &logger{
			logger: logrus.New(),
		}
	})
	return log
}

//
func (l *logger) Info(message string) {
	if Config().Debug {
		l.logger.Info(message)
	}
}

//
func (l *logger) Print(message string) {
	if Config().Debug {
		l.logger.Print(message)
	}
}

//
func (l *logger) Println(message string) {
	if Config().Debug {
		l.logger.Println(message)
	}
}
