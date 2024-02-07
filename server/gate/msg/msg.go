package msg

import (
	"message"

	"github.com/name5566/leaf/network/json"
)

var JSONProcessor = json.NewProcessor()

func init() {
	JSONProcessor.Register(&message.Login{})
}
