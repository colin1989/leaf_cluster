package world

import (
	"message"
	"server/world/msg"
)

func init() {
	msg.JSONProcessor.SetRouter(&message.S2S_Reg{}, ChanRPC)
}
