package internal

import (
	"fmt"
	"message"
	"server/gate/constant"

	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	//skeleton.RegisterChanRPC(constant.NewWorldFunc, NewWorldFunc)
	skeleton.RegisterChanRPC(constant.NewGameFunc, NewGameFunc)
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func NewWorldFunc(args []interface{}) {
	a := args[0].(gate.Agent)
	//a.SetUserData(&GameData{})

	a.WriteMsg(&message.S2S_Reg{Server: &protos.Server{
		ID:      serverID,
		Address: wsAddr,
		Typ:     protos.ServerType_Gate,
	}})
	fmt.Println("rpcNewWorldFunc!!!")
}

func NewGameFunc(args []interface{}) {
	a := args[0].(gate.Agent)
	sid := args[1].(int32)

	_, ok := GameAgents[sid]
	if ok {
		log.Error("重复连接游戏服", log.Int32("sid", sid))
	}
	GameAgents[sid] = a
	fmt.Println("rpcNewGameServer!!!")
}

var agentId = 0

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	agentId++
	a.SetUserData(&AgentData{
		agentId: agentId,
	})
	AgentMap[agentId] = a

	fmt.Println("rpcNewAgent!!!")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	switch data := a.UserData().(type) {
	case *AgentData:
		if data.agentId == 0 {
			return
		}
		delete(AgentMap, data.agentId)
		SendToGame(data.serverID, &message.Disconnect{UserID: data.userID})
	}
	fmt.Println("rpcCloseAgent!!!")
}
