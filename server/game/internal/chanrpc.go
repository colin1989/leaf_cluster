package internal

import (
	"fmt"
	"message"
	"server/protos"

	"github.com/name5566/leaf/gate"
)

const NewWorldAgentFunc = "NewWorldFunc"

func init() {
	skeleton.RegisterChanRPC(NewWorldAgentFunc, NewWorldFunc)
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func NewWorldFunc(args []interface{}) {
	a := args[0].(gate.Agent)
	//a.SetUserData(&GameData{})

	a.WriteMsg(&message.S2S_Reg{Server: &protos.Server{
		ID:      serverID,
		Address: wsAddr,
		Typ:     protos.ServerType_Game,
	}})
	fmt.Println("rpcNewWorldFunc!!!")
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	//a.WriteMsg()

	fmt.Println("rpcNewAgent!!!")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	a.Close()

	fmt.Println("rpcCloseAgent!!!")
}
