package handlers

import (
	"net/http"

	"github.com/francky-d/go-forum/internal/utils/helpers"
)

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	helpers.RenderView("home", nil, response)
}

func RedirectToHomeHandler(response http.ResponseWriter, request *http.Request) {
	http.Redirect(response, request, "/home", http.StatusMovedPermanently)
}
