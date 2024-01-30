package msg

import (
	"github.com/name5566/leaf/network/json"
	"message"
)

var JSONProcessor = json.NewProcessor()

func init() {

	JSONProcessor.Register(&message.S2S_Msg{})
	JSONProcessor.Register(&message.S2S_Reg{})
	JSONProcessor.Register(&message.C2S_Msg{})
	JSONProcessor.Register(&message.S2C_Msg{})

	JSONProcessor.Register(&message.Login{})
	JSONProcessor.Register(&message.Bind{})
}
