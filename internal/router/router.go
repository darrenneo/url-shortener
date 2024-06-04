package router

import (
	"net/http"

	"url-shortener/internal/handler"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPost},
		}),
		echoMiddleware.Gzip(),
		echoMiddleware.Secure(),
	)

	attachRoutes(e)
	return e
}

func attachRoutes(e *echo.Echo) {
	e.POST("/shorten", handler.ShortenURL)
	e.GET("/:url", handler.RedirectURL)
	e.GET("/debug/list", handler.ListAll)
}
