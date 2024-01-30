package msg

import (
	"message"

	"github.com/name5566/leaf/network/json"
)

var JSONProcessor = json.NewProcessor()

func init() {
	JSONProcessor.Register(&message.S2S_Reg{})
	JSONProcessor.Register(&message.W2S_GS{})
}
