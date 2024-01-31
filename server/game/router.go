package game

import (
	"message"
	"server/game/msg"
)

func init() {
	msg.JSONProcessor.SetRouter(&message.Gate_Forward{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.Disconnect{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.Kick{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.Login{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.Greeting{}, ChanRPC)
}
