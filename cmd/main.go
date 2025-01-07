package main

import (
	"github.com/francky-d/go-forum/internal/utils"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", utils.Router()))
}
