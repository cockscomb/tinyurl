package usecase

import "github.com/google/wire"

var UsecaseSet = wire.NewSet(
	NewTinyURLUsecase,
)
