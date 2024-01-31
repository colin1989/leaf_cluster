package msg

import (
	"message"

	"github.com/name5566/leaf/network/json"
)

var JSONProcessor = json.NewProcessor()

func init() {
	JSONProcessor.Register(&message.S2W_GS{})
	JSONProcessor.Register(&message.W2S_GS{})
	JSONProcessor.Register(&message.Gate_Forward{})
	JSONProcessor.Register(&message.S2S_Reg{})
	JSONProcessor.Register(&message.C2S_Msg{})
	JSONProcessor.Register(&message.S2C_Msg{})
	JSONProcessor.Register(&message.Disconnect{})
	JSONProcessor.Register(&message.Kick{})

	JSONProcessor.Register(&message.Login{})
	JSONProcessor.Register(&message.Bind{})
}
