package entity

import (
	"github.com/labstack/gommon/random"
	"net/url"
)

type TinyURL struct {
	ID  string
	URL *url.URL
}

func GenerateTinyURL(URL *url.URL) *TinyURL {
	return &TinyURL{
		ID:  random.String(8, random.Alphanumeric),
		URL: URL,
	}
}
