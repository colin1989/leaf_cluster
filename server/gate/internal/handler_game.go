package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"message"
)

func init() {
	// rpc
	handleMsg(&message.S2S_Reg{}, S2S_Reg)
	handleMsg(&message.Bind{}, Bind)
}

func S2S_Reg(args []interface{}) {
	m := args[0].(*message.S2S_Reg)
	a := args[1].(gate.Agent) // game server
	if a == nil || m == nil {
		fmt.Println("greeting err")
		return
	}

	a.UserData().(*GameData).serverID = m.Id
	GameServers[m.Id] = a
}

func Bind(args []interface{}) {
	m := args[0].(*message.Bind)
	//a := args[1].(gate.Agent) // game server
	//if a == nil || m == nil {
	//	log.Infof("greeting err")
	//	return
	//}

	agent, ok := AgentMap[m.Agent]
	if !ok {
		log.Errorf("wrong agent id %v", m.Agent)
		return
	}

	agentData := agent.UserData().(*AgentData)
	if agentData == nil {
		log.Infof("agentData == nil")
		agent.Close()
		return
	}

	if agentData.serverID != 0 {
		log.Infof("agent has already logged int ")
		agent.Close()
		return
	}

	agentData.serverID = m.Server
	agentData.userID = m.UserID
	agent.WriteMsg(&message.Login{
		Server:  m.Server,
		Account: agentData.accID,
		Agent:   m.Agent,
	})
}
