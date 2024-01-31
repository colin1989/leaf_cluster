package internal

import (
	"message"
	"reflect"
	"server/gate/constant"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&message.W2S_GS{}, W2S_GS)
	handleMsg(&message.C2S_Msg{}, C2S_Msg)
	handleMsg(&message.Login{}, Login)
}

// W2S_GS
//
//	@Description: world推送游戏服信息
//	@param args
func W2S_GS(args []interface{}) {
	m := args[0].(*message.W2S_GS)
	//a := args[1].(gate.Agent)

	for _, server := range m.Servers {
		// TODO 这个地方可能会有问题
		if client, ok := GameClients[server.ID]; ok {
			// Close 里面有个wait...
			go client.Close()
		}
		if _, ok := GameAgents[server.ID]; ok {
			GameAgents[server.ID].Close()
			delete(GameAgents, server.ID)
		}
		GameClients[server.ID] = Gate.Connect(server.ID, server.Address, constant.NewGameFunc, 1)
	}
}

//func Gate_Forward(args []interface{}) {
//	m := args[0].(*message.Gate_Forward)
//	a := args[1].(gate.Agent) // game server
//	if a == nil || m == nil {
//		log.Infof("greeting err")
//		return
//	}
//
//	b := m.Body
//	var greeting message.Greeting
//	_ = json.Unmarshal(b, &greeting)
//
//	log.Infof("from", m.From, "to", m.To, greeting.Code, greeting.Message)
//
//	// 转发给客户端
//	if Client != nil {
//		log.Infof("send to client")
//		Client.WriteMsg(&greeting)
//	}
//}

func C2S_Msg(args []interface{}) {
	m := args[0].(*message.C2S_Msg)
	a := args[1].(gate.Agent) // game server
	if a == nil || m == nil {
		log.Infof("greeting err")
		return
	}

	agentData := a.UserData().(*AgentData)
	if agentData == nil {
		log.Infof("agentData == nil")
		a.Close()
		return
	}

	if agentData.serverID == 0 {
		log.Infof("the first message must to be LOGIN")
		a.Close()
		return
	}

	msg := &message.Gate_Forward{
		From:   "",
		To:     "",
		MsgID:  "",
		Agent:  agentData.agentId,
		UserID: agentData.userID,
		Body:   m.Body,
	}

	if err := SendToGame(agentData.serverID, msg); err != nil {
		log.Infof("SendToGame error : ", agentData.serverID)
		a.Close()
		return
	}
}

func Login(args []interface{}) {
	m := args[0].(*message.Login)
	a := args[1].(gate.Agent) // game server
	if a == nil || m == nil {
		log.Infof("greeting err")
		return
	}

	agentData := a.UserData().(*AgentData)
	if agentData == nil {
		log.Infof("agentData == nil")
		a.Close()
		return
	}

	if agentData.serverID != 0 {
		log.Infof("agent has already logged int ")
		a.Close()
		return
	}

	m.Agent = agentData.agentId

	if err := SendToGame(m.Server, m); err != nil {
		log.Infof("SendToGame error : ", agentData.serverID)
		a.Close()
		return
	}
}
