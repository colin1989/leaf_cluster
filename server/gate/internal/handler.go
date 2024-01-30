package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"message"
	"reflect"
	"server/constant"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&message.C2S_Msg{}, C2S_Msg)
	handleMsg(&message.S2C_Msg{}, S2C_Msg)
	handleMsg(&message.Login{}, Login)
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

	if agentData.serverID == 0 && m.MsgID != constant.LOGIN {
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
	server.WriteMsg(m)
}

func S2C_Msg(args []interface{}) {
	m := args[0].(*message.S2C_Msg)
	a := args[1].(gate.Agent) // game server
	if a == nil || m == nil {
		fmt.Println("greeting err")
		return
	}

	agent, ok := AgentMap[m.Agent]
	if !ok {
		log.Errorf("wrong agent id %v", m.Agent)
		return
	}
	agent.WriteRaw(m.MsgID, m.Body)
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
