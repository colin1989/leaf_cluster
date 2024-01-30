package world

import (
	"server/world/internal"
)

var (
	Module     = new(internal.Module)
	GateModule = new(internal.GateModule)
	ChanRPC    = internal.ChanRPC
)
