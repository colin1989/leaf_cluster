package game

import (
	"server/game/internal"
	"server/msg"
)

var (
	Module     = new(internal.Module)
	GateModule = new(internal.GateModule)
	ChanRPC    = internal.ChanRPC
)

func init() {
	msg.JSONProcessor.SetRouter(&msg.S2S_Msg{}, ChanRPC)
}
