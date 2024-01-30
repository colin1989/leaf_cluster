package game

import (
	"server/game/internal"
)

var (
	Module     = new(internal.Module)
	GateModule = new(internal.GateModule)
	ChanRPC    = internal.ChanRPC
)

func init() {
	//msg.JSONProcessor.SetRouter(&message.S2S_Msg{}, ChanRPC)
}

func SetServerID(id int32) {
	internal.SetServerID(id)
}

func SetWSAddr(addr string) {
	internal.SetWSAddr(addr)
}
