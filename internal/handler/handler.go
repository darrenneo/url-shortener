package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

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
	return c.JSON(http.StatusOK, "URL shortened")
}

func RedirectURL(c echo.Context) error {
	url := c.Param("url")
	originalURL, exist := storage.GetURL(url)

	if !exist {
		return c.JSON(http.StatusNotFound, "URL not found")
	}
	redirectUrl := checkPrefix(originalURL)
	return c.Redirect(http.StatusMovedPermanently, redirectUrl)
}

func checkPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	return url
}

func ListAll(c echo.Context) error {
	storage.ListAll()
	return c.JSON(http.StatusOK, "Listed all URLs")
}
