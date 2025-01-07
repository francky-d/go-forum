package main

import "github.com/francky-d/go-forum/internal/utils/helpers"

func main() {
	helpers.RenderViewToStdout("home.gohtml", nil)
}
