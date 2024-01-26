package msg

import "github.com/name5566/leaf/network/json"

var JSONProcessor = json.NewProcessor()

func init() {

	JSONProcessor.Register(&Greeting{})

}

type (
	Greeting struct {
		Code    int
		Message string
	}
)
