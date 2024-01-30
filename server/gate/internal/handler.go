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
		if _, ok := GameClients[server.ID]; ok {
			GameClients[server.ID].Close()
		}
		if _, ok := GameServers[server.ID]; ok {
			GameServers[server.ID].Close()
			delete(GameServers, server.ID)
		}
		GameClients[server.ID] = Gate.Connect(server.ID, server.Address, constant.NewGameFunc, 1)
	}
}

//func S2S_Msg(args []interface{}) {
//	m := args[0].(*message.S2S_Msg)
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

	if agentData.serverID == 0 && m.MsgID != "login" {
		log.Infof("the first message must to be LOGIN")
		a.Close()
		return
	}

	server, ok := GameServers[agentData.serverID]
	if !ok {
		log.Infof("log in wrong server : ", agentData.serverID)
		a.Close()
		return
	}

	//b, _ := json.Marshal(&m)
	m.Agent = agentData.agentId
	m.UserID = agentData.userID
	server.WriteMsg(m.Body)
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

	server, ok := GameServers[m.Server]
	if !ok {
		log.Infof("log in wrong server : ", m.Server)
		a.Close()
		return
	}

	//b, _ := json.Marshal(&m)
	//server.WriteMsg(&message.S2S_Msg{
	//	From:  "gate",
	//	To:    "game",
	//	MsgID: message.GetMsgId(m),
	//	Body:  b,
	//})
	m.Agent = agentData.agentId
	server.WriteMsg(m)
}
