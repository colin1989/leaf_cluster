package game

import (
	"server/game/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func SetServerID(id int32) {
	internal.SetServerID(id)
}

func SetWSAddr(addr string) {
	internal.SetWSAddr(addr)
}

func init() {
}
