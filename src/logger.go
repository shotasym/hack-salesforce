package main

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(logrus.InfoLevel.String())
	if err == nil {
		logger.Level = level
	}

	return logger
}
