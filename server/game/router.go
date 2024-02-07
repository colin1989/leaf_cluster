package game

import (
	"message"
	"server/game/msg"
)

func init() {
	msg.JSONProcessor.SetRouter(&message.Login{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.Greeting{}, ChanRPC)
}
