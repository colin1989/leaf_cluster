package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"message"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	reg := new(message.S2S_Reg)
	reg.Key = Key
	reg.Id = serverID
	reg.Ip = wsAddr

	a.WriteMsg(reg)

	fmt.Println("rpcNewAgent!!!")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	a.Close()

	fmt.Println("rpcCloseAgent!!!")
}
