package main

import (
	"log"
	"net/http"

	"github.com/odmrs/trakrlist-api/internal/router"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", router.Ping)

	log.Print("Listening...")
	http.ListenAndServe(":8080", mux)
}
