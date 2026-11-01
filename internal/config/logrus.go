package config

import (
	"strconv"

	"github.com/sirupsen/logrus"
)

func NewLogger(config *ConfigEnv) *logrus.Logger {
	log := logrus.New()

	level, err := strconv.Atoi(config.LOG_LEVEL)
	if err != nil {
		level = 1
	}

	log.SetLevel(logrus.Level(level))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
