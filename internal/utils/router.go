package utils

import (
	"net/http"

	"github.com/francky-d/go-forum/internal/handlers"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /", handlers.RedirectToHomeHandler)

	router.HandleFunc("GET /home", handlers.HomeHandler)
	return router
}
