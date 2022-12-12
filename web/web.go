package web

import "github.com/google/wire"

var WebSet = wire.NewSet(
	NewServer,
	NewTemplate,
)
