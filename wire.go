//go:build wireinject

package main

import (
	"context"
	"github.com/cockscomb/tinyurl/web"
	"github.com/google/wire"
)

func InitializeServer(ctx context.Context) (*web.Server, error) {
	wire.Build(
		ConfigSet,
		AWSSet,
		web.NewServer,
	)
	return &web.Server{}, nil
}
