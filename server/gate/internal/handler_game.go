package internal

import (
	"message"

	"github.com/name5566/leaf/log"
)

func init() {
	// rpc
	handleMsg(&message.Kick{}, Kick)
	handleMsg(&message.S2C_Msg{}, S2C_Msg)
	handleMsg(&message.Bind{}, Bind)
}

func S2C_Msg(args []interface{}) {
	m := args[0].(*message.S2C_Msg)
	//a := args[1].(gate.Agent) // game server

	agent, ok := AgentMap[m.Agent]
	if !ok {
		log.Errorf("wrong agent id %v", m.Agent)
		return
	}
	agent.WriteRaw(m.Body)
}

func Kick(args []interface{}) {
	m := args[0].(*message.Kick)

	agent, ok := AgentMap[m.Agent]
	if !ok {
		log.Errorf("wrong agent id %v", m.Agent)
		return
	}
	agent.Close()
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
