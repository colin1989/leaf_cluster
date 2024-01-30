package internal

import (
	"fmt"
	"message"
	"server/gate/constant"
	"server/protos"

	"github.com/name5566/leaf/log"

	"github.com/name5566/leaf/network"

	"github.com/name5566/leaf/gate"
)

type AgentData struct {
	accID    int
	userID   int
	serverID int32
	agentId  int
}

var GameClients = map[int32]*network.WSClient{}
var GameServers = map[int32]gate.Agent{}
var AgentMap = map[int]gate.Agent{}

func init() {
	skeleton.RegisterChanRPC(constant.NewWorldFunc, NewWorldFunc)
	skeleton.RegisterChanRPC(constant.NewGameFunc, NewGameFunc)
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func NewWorldFunc(args []interface{}) {
	a := args[0].(gate.Agent)
	//a.SetUserData(&GameData{})

	a.WriteMsg(&message.S2S_Reg{Server: &protos.Server{
		ID:      1000,
		Address: "127.0.0.1:13563",
		Typ:     protos.ServerType_Gate,
	}})
	fmt.Println("rpcNewWorldFunc!!!")
}

func NewGameFunc(args []interface{}) {
	a := args[0].(gate.Agent)
	sid := args[1].(int32)

	_, ok := GameServers[sid]
	if ok {
		log.Error("重复连接游戏服", log.Int32("sid", sid))
	}
	GameServers[sid] = a
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
	}
	fmt.Println("rpcCloseAgent!!!")
}
