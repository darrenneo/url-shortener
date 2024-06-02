package router

import (
	"url-shortener/internal/handler"

	"github.com/labstack/echo/v4"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	attachRoutes(e)
	return e
}

func attachRoutes(e *echo.Echo) {
	e.POST("/shorten", handler.ShortenURL)
	e.GET("/:url", handler.RedirectURL)
	e.GET("/debug/list", handler.ListAll)
}
