package internal

import (
	"fmt"

	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	fmt.Println("rpcNewAgent!!!")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	s, ok := a.UserData().(*protos.Server)
	if ok {
		manager.RemoveServer(s.Typ, s.ID)
	}

	a.Close()

	fmt.Println("rpcCloseAgent!!!")
}
