package game

import (
	"message"
	"server/msg"
)

func init() {
	msg.JSONProcessor.Register(&message.Greeting{})
}

func init() {
	msg.JSONProcessor.SetRouter(&message.Login{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.C2S_Msg{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.Greeting{}, ChanRPC)
}
