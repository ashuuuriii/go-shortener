package models

type ShortenRequest struct {
	LongURL string `json:"long_url" validate:"required,url"`
}

type ShortenResponse struct {
	ShortenedUrl string `json:"shortened_url"`
}
