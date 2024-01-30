package gate

import "server/gate/internal"

var (
	Module     = new(internal.GateModule)
	GateModule = new(internal.Module)
	ChanRPC    = internal.ChanRPC
)
