package main

import (
	"github.com/eliabedossantos/gopportunities/config"
	"github.com/eliabedossantos/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	println("initializing gopportunities")

	//initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialiize Router
	router.Initialize()
}
