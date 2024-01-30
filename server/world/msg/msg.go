package msg

import (
	"github.com/name5566/leaf/network/json"
	"message"
)

var JSONProcessor = json.NewProcessor()

func init() {
	JSONProcessor.Register(&message.S2S_Reg{})
}
