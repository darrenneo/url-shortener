package handler

import "github.com/labstack/echo/v4"

func ShortenURL(c echo.Context) error {
	url := c.Param("url")
	return c.String(200, url)
}

func RedirectURL(c echo.Context) error {
	return nil
}
