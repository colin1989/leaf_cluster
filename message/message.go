package message

import "reflect"

type (
	// 测试
	C2S_Gates struct {
	}

	// 测试
	S2C_Gates struct {
		GameID    []int32
		Addresses []string
	}

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
