package main

import (
	"os"
	"os/user"
	"strconv"

	"github.com/sirupsen/logrus"
)

func initLogger() {
	var debug bool
	s := os.Getenv(DebugEnvVar)
	if s == "" {
		debug = false
	} else {
		debug, _ = strconv.ParseBool(s)
	}
	if debug {
		logrus.SetOutput(os.Stdout)
	} else {
		lesenDir := homeDir() + DefaultLesenDir
		os.Mkdir(lesenDir, os.ModePerm)
		logPath := lesenDir + DefaultLogFile
		logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(logFile)
	}
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}
