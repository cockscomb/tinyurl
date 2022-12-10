//go:build wireinject

package main

import (
	"github.com/cockscomb/tinyurl/web"
	"github.com/google/wire"
)

func InitializeServer() (*web.Server, error) {
	wire.Build(
		ConfigSet,
		web.NewServer,
	)
	return &web.Server{}, nil
}
