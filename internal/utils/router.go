package utils

import (
	"github.com/francky-d/go-forum/internal/utils/helpers"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(response http.ResponseWriter, request *http.Request) {
		http.Redirect(response, request, "/home", http.StatusMovedPermanently)
	})

	router.HandleFunc("GET /home", func(response http.ResponseWriter, request *http.Request) {
		helpers.RenderView("home", nil, response)
	})
	return router
}
