package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

var GameSvr gate.Agent

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	// 连接到game server
	if GameSvr == nil {
		//TODO 通过 manager 管理 game server 的连接
		GameSvr = a
	}

	fmt.Println("rpcNewAgent!!!")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	// 断开连接game server
	if GameSvr != nil {
		GameSvr = nil
	}

	fmt.Println("rpcCloseAgent!!!")
}
