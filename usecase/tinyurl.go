package usecase

import (
	"context"
	"errors"
	"github.com/cockscomb/tinyurl/domain/entity"
	"net/url"
)

type URLStore interface {
	Create(ctx context.Context, url *entity.TinyURL) error
	Find(ctx context.Context, id string) (*entity.TinyURL, error)
}

type TinyURLUsecase struct {
	store URLStore
}

func NewTinyURLUsecase(store URLStore) *TinyURLUsecase {
	return &TinyURLUsecase{store: store}
}

func (usecase *TinyURLUsecase) Generate(ctx context.Context, u *url.URL) (*entity.TinyURL, error) {
	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, entity.NewValidationError("invalid scheme")
	}
	for {
		tu := entity.GenerateTinyURL(u)
		if err := usecase.store.Create(ctx, tu); err != nil {
			if errors.Is(err, entity.ErrAlreadyExists) {
				continue
			}
			return nil, err
		}
		return tu, nil
	}
}

func (usecase *TinyURLUsecase) Access(ctx context.Context, id string) (*entity.TinyURL, error) {
	return usecase.store.Find(ctx, id)
}

func (usecase *TinyURLUsecase) Peek(ctx context.Context, id string) (*entity.TinyURL, error) {
	return usecase.store.Find(ctx, id)
}
