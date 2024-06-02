package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"url-shortener/internal/model"
	"url-shortener/internal/storage"

	"github.com/labstack/echo/v4"
)

func ShortenURL(c echo.Context) error {
	if c.Request().Body == nil {
		return c.JSON(http.StatusBadRequest, "Missing body")
	}

	rawData, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid body")
	}

	apiRequest := model.ApiRequest{}
	json.Unmarshal(rawData, &apiRequest)

	err = storage.AddURL(apiRequest.OriginalURL, apiRequest.NewURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	storage.ListAll()

	return nil
}

func RedirectURL(c echo.Context) error {
	return nil
}
