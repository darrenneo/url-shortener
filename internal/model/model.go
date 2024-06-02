package model

type ApiRequest struct {
	OriginalURL string `json:"original_url"`
	NewURL      string `json:"new_url"`
}
