package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/AkihikoITOH/lesen/backend"
	"github.com/AkihikoITOH/lesen/config"
	"github.com/AkihikoITOH/lesen/util"
)

func initLogger() *logrus.Logger {
	os.Mkdir(util.LesenDir(), os.ModePerm)
	logFile, err := os.OpenFile(util.LogPath(), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetOutput(logFile)

	return logger
}

func initConfig() {
	logger := initLogger()

	config.SetLogger(logger)
	config.SetBackend(backend.NewDiskvBackend(util.DiskvDir()))
}
