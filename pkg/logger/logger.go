package logger

import "github.com/juju/loggo"

var logger = loggo.GetLogger(loggo.DefaultWriterName)

func Info(s string, args... any) {
	logger.Infof(s, args...)
}