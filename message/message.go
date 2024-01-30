package message

import "reflect"

type (
	S2S_Msg struct {
		From  string
		To    string
		MsgID string
		Body  []byte
	}
	//连接注册
	S2S_Reg struct {
		Key string //密钥
		Id  int    //服务器ID
		Ip  string //地址
	}

	C2S_Msg struct {
		From   string
		To     string
		MsgID  string
		Agent  int
		UserID int
		Body   []byte
	}

	S2C_Msg struct {
		From  string
		To    string
		MsgID string
		Agent int
		Body  []byte
	}
)

type (
	Login struct {
		Server  int
		Account int
		Agent   int
	}

	Bind struct {
		Agent  int
		Server int
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
