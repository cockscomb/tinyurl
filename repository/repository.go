package repository

import (
	"github.com/cockscomb/tinyurl/usecase"
	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	NewURLStore,
	wire.Bind(new(usecase.URLStore), new(*URLStore)),
)
