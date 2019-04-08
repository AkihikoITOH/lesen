package config

import (
	"github.com/sirupsen/logrus"

	"github.com/AkihikoITOH/lesen/backend"
)

var instance = &config{}

type config struct {
	backend backend.Backend
	logger  *logrus.Logger
}

func Backend() backend.Backend {
	return instance.backend
}

func SetBackend(be backend.Backend) {
	instance.backend = be
}

func Logger() *logrus.Logger {
	return instance.logger
}

func SetLogger(lgr *logrus.Logger) {
	instance.logger = lgr
}
