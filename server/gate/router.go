package gate

import (
	"message"
	"server/gate/msg"
)

func init() {
	msg.JSONProcessor.SetRouter(&message.Login{}, ChanRPC)
}
