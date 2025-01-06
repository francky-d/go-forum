package main

import "github.com/francky-d/go-forum/internal/utils"

func main() {
	utils.RenderViewToStdout("home.gohtml", nil)
}
