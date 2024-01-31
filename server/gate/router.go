package gate

import (
	"message"
	"server/gate/msg"
)

func init() {
	msg.JSONProcessor.SetRouter(&message.W2S_GS{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.C2S_Msg{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.S2C_Msg{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.Kick{}, ChanRPC)
	// 消息路由到Game server
	msg.JSONProcessor.SetRouter(&message.Login{}, ChanRPC)
	//msg.JSONProcessor.SetRouter(&message.Greeting{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.Bind{}, ChanRPC)
}
