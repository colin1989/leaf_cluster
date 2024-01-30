package message

import "server/protos"

const (
	ServerTypeGate = 1
	ServerTypeGame = 1
)

// world
type (
	//连接注册
	S2S_Reg struct {
		Server *protos.Server
	}

	S2W_GS struct {
	}

	W2S_GS struct {
		Servers []*protos.Server
	}
)

type (
	S2S_Msg struct {
		From  string
		To    string
		MsgID string
		Body  []byte
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
