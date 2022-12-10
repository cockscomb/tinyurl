package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/cockscomb/tinyurl/web"
	"github.com/google/wire"
)

type Config struct {
	web.ServerConfig
}

func ParseEnv() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

var ConfigSet = wire.NewSet(
	ParseEnv,
	wire.FieldsOf(new(*Config), "ServerConfig"),
)
