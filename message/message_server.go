package message

import "server/protos"

// world
type (
	// S2S_Reg 非 world 服务启动时，主动向 world 注册服务信息
	S2S_Reg struct {
		Server *protos.Server
	}

	// S2W_GS 请求游戏服数据
	S2W_GS struct {
	}

	// W2S_GS 通知游戏服数据
	W2S_GS struct {
		Servers []*protos.Server
	}
)

type (
	// Gate_Forward 网关封装客户端数据转发游戏服
	// 该消息的前置为 C2S_Msg
	Gate_Forward struct {
		From   string
		To     string
		MsgID  string
		Agent  int
		UserID int
		Body   []byte
	}

	// C2S_Msg 网关服收到未能反序列化的消息时，进行封装处理
	// 之后转成 Gate_Forward
	C2S_Msg struct {
		Body []byte
	}

	// S2C_Msg 游戏服转发至网关服的消息
	S2C_Msg struct {
		From  string
		To    string
		MsgID string
		Agent int
		Body  []byte
	}

	// Kick 游戏服收到 Kick 后，转发至网关
	Kick struct {
		UserID int
		Agent  int
	}

	// Disconnect 网关服通知游戏服断开连接
	Disconnect struct {
		UserID int
	}
)
