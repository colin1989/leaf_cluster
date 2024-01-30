package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
)

type GameData struct {
	serverID int
}

type AgentData struct {
	accID    int
	userID   int
	serverID int
	agentId  int
}

var GameServers = map[int]gate.Agent{}
var AgentMap = map[int]gate.Agent{}

func init() {
	skeleton.RegisterChanRPC("NewGameServer", NewGameServer)
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func NewGameServer(args []interface{}) {
	a := args[0].(gate.Agent)
	a.SetUserData(&GameData{})

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
	case *GameData:
		if data.serverID == 0 {
			return
		}
		delete(GameServers, data.serverID)
	case *AgentData:
		if data.agentId == 0 {
			return
		}
		delete(AgentMap, data.agentId)
	}
	fmt.Println("rpcCloseAgent!!!")
}
