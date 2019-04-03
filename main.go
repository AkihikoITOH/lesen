package main

import (
	"github.com/sirupsen/logrus"

	"github.com/AkihikoITOH/lesen/cmd"
)

const (
	DebugEnvVar     = "LESEN_DEBUG_MODE"
	DefaultLesenDir = "/.lesen"
	DefaultLogFile  = "/lesen.log"
)

func main() {
	initLogger()
	err := cmd.StartNewCLIApp()

	if err != nil {
		logrus.Fatal(err.Error())
	}
}
