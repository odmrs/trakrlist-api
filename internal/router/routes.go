package router

import (
	"net/http"

	"github.com/odmrs/trakrlist-api/internal/handler"
)

func initializeRoutes(v1 *http.ServeMux) {
	// Initilize Handlers
	handler.InitializeHandler()
	v1.HandleFunc("GET /api/v1/ping", handler.Ping)

	// Content handlers
	v1.HandleFunc("GET /api/v1/content/{content}", handler.ShowContentHandler)
	v1.HandleFunc("POST /api/v1/content", handler.CreateContentHandler)
	v1.HandleFunc("PUT /api/v1/content", handler.Ping)
	v1.HandleFunc("DELETE /api/v1/content/{content}", handler.DeleteContentHandler)
	v1.HandleFunc("GET /api/v1/contents", handler.ListAllContentHandler)
}
