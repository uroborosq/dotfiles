package logger

import (
	"os"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger(loggo.DefaultWriterName)

func Infof(s string, args ...any) {
	logger.Infof(s, args...)
}

func Warnf(s string, args ...any) {
	logger.Warningf(s, args...)
}

func Fatalf(s string, args ...any) {
	logger.Errorf(s, args...)
	os.Exit(1)
}

func Errorf(s string, args ...any) {
	logger.Errorf(s, args...)
}
