package b

import (
	"github.com/google/wire"
	"wire_best_practice/a"
)

var ProviderB = wire.NewSet(NewB, a.ProviderA, wire.Bind(new(BInterface), new(*B)))
