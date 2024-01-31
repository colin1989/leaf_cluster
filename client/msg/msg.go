package msg

import (
	"message"

	"github.com/name5566/leaf/network/json"
)

var JSONProcessor = json.NewProcessor()

func init() {

	JSONProcessor.Register(&message.S2C_Gates{})
	JSONProcessor.Register(&message.C2S_Gates{})
	JSONProcessor.Register(&message.Login{})
	JSONProcessor.Register(&message.Greeting{})

}
