package entity

import "net/url"

type TinyURL struct {
	ID  string
	URL *url.URL
}
