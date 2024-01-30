package message

import "reflect"

type (
	Login struct {
		Server  int32
		Account int
		Agent   int
	}

	Bind struct {
		Agent  int
		Server int32
		UserID int
	}

	Greeting struct {
		Code    int
		Message string
	}
)

func GetMsgId(msg interface{}) string {
	return reflect.TypeOf(msg).Elem().Name()
}
