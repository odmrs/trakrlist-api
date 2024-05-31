package router

import (
	"log"
	"net/http"
)

func Initialize() {
	v1 := http.NewServeMux()

	initializeRoutes(v1)

	server := &http.Server{Addr: "localhost:8080", Handler: v1}

	log.Print("Listening...")
	server.ListenAndServe()
}
