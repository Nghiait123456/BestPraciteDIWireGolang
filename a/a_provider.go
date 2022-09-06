package a

import "github.com/google/wire"

var ProviderA = wire.NewSet(
	NewA2, wire.Bind(new(AInterface), new(*A2)),
)
