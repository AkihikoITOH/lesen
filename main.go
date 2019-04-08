package main

import (
	"github.com/AkihikoITOH/lesen/cmd"
	"github.com/AkihikoITOH/lesen/config"
)

func main() {
	initConfig()

	err := cmd.StartNewCLIApp()

	if err != nil {
		config.Logger().Fatal(err.Error())
	}
}
