package msg

import (
	"github.com/name5566/leaf/network/json"
	"reflect"
)

var JSONProcessor = json.NewProcessor()

func init() {

	JSONProcessor.Register(&S2S_Msg{})
	JSONProcessor.Register(&Greeting{})

}

type (
	C2S_Msg struct {
		MsgID   string
		MsgType interface{}
		UID     int
		Body    []byte
	}

	S2S_Msg struct {
		From  string
		To    string
		MsgID string
		UID   int
		Body  []byte
	}

	Greeting struct {
		Code    int
		Message string
	}
)

func GetMsgId(msg interface{}) string {
	return reflect.TypeOf(msg).Elem().Name()
}
