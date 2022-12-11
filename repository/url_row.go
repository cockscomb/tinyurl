package repository

import (
	"github.com/cockscomb/tinyurl/domain/entity"
	"net/url"
	"time"
)

type URLRow struct {
	ID        string    `dynamodbav:"id"`
	URL       string    `dynamodbav:"url"`
	CreatedAt time.Time `dynamodbav:"created_at"`
}

func (r URLRow) ToEntity() (*entity.TinyURL, error) {
	u, err := url.Parse(r.URL)
	if err != nil {
		return nil, err
	}
	return &entity.TinyURL{
		ID:  r.ID,
		URL: u,
	}, nil
}
