package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func SetOutput(w io.Writer) {
	logger.Out = w
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func WithError(err error) *logrus.Entry {
	return logger.WithField("error", err)
}

func SetLevel(level logrus.Level) {
	logger.SetLevel(level)
}

type Fields = logrus.Fields

var DebugLevel = logrus.DebugLevel
