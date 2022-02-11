package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger() error {
	if logger == nil {
		logger = logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetLevel(logrus.DebugLevel)
		logger.SetOutput(os.Stdout)
	}

	return nil
}

func Get() *logrus.Logger {
	return logger
}
