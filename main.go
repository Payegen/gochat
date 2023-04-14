package main

import "gochat/router"

func main() {
	r := router.Router()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.Run(":8081")
}
