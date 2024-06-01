package router

import (
	"net/http"

	"github.com/odmrs/trakrlist-api/config"
)

var logger config.Logger

func Initialize() {
	logger = *config.GetLogger("router")
	v1 := http.NewServeMux()

	initializeRoutes(v1)

	server := &http.Server{Addr: "localhost:8080", Handler: v1}

	logger.Info("Listening...")
	server.ListenAndServe()
}
