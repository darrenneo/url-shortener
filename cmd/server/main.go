package main

import (
	"url-shortener/internal/router"
)

func main() {
	e := router.InitRouter()
	e.Logger.Fatal(e.Start("localhost:8000"))
}
