package gate

import (
	"server/gate/internal"
)

var (
	Module     = new(internal.Module)
	GateModule = new(internal.GateModule)
	ChanRPC    = internal.ChanRPC
)
