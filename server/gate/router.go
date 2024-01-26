package gate

import (
	"server/msg"
)

func init() {
	// 消息路由到Game server
	msg.JSONProcessor.SetRouter(&msg.Greeting{}, ChanRPC)

	msg.JSONProcessor.SetRouter(&msg.S2S_Msg{}, ChanRPC)
}
