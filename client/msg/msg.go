package msg

import (
	"github.com/name5566/leaf/network/json"
	"message"
)

var JSONProcessor = json.NewProcessor()

func init() {

	JSONProcessor.Register(&message.Login{})
	JSONProcessor.Register(&message.Greeting{})

}
