//go:build wireinject

package main

import (
	"context"
	"github.com/cockscomb/tinyurl/repository"
	"github.com/cockscomb/tinyurl/usecase"
	"github.com/cockscomb/tinyurl/web"
	"github.com/cockscomb/tinyurl/web/controller"
	"github.com/google/wire"
)

func InitializeServer(ctx context.Context) (*web.Server, error) {
	wire.Build(
		ConfigSet,
		AWSSet,
		repository.RepositorySet,
		usecase.UsecaseSet,
		controller.NewController,
		web.WebSet,
	)
	return &web.Server{}, nil
}
