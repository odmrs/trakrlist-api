package main

import (
	"github.com/odmrs/trakrlist-api/config"
	"github.com/odmrs/trakrlist-api/internal/router"
)

var logger config.Logger

func main() {
	logger = *config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
