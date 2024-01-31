package internal

import (
	"fmt"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/network"
)

type AgentData struct {
	accID    int
	userID   int
	serverID int32
	agentId  int
}

var GameClients = map[int32]*network.WSClient{}
var GameAgents = map[int32]gate.Agent{}
var AgentMap = map[int]gate.Agent{}

func SendToGame(sid int32, msg interface{}) error {
	server, ok := GameAgents[sid]
	if !ok {
		return fmt.Errorf("wrong server : %v", sid)
	}
	server.WriteMsg(msg)
	return nil
}
