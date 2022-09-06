package c

import (
	"github.com/google/wire"
	"wire_best_practice/b"
)

var ProviderC = wire.NewSet(NewC, b.ProviderB, wire.Bind(new(CInterface), new(*C)))
