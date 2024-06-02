package main

import (
	"url-shortener/internal/router"
	"url-shortener/internal/storage"
)

func main() {
	e := router.InitRouter()
	storage.InitStorage()
	e.Logger.Fatal(e.Start("localhost:8000"))
}
