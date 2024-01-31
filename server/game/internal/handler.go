package internal

import (
	"math/rand"
	"message"
	"reflect"
	"server/game/msg"

	"github.com/name5566/leaf/log"

	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&message.Gate_Forward{}, Gate_Forward)
	handleMsg(&message.Disconnect{}, Disconnect)
	handleMsg(&message.Kick{}, Kick)
	handleMsg(&message.Login{}, Login)
	handleMsg(&message.Greeting{}, Greeting)
}

// Gate_Forward
//
//	@Description:
//	@param args
func Gate_Forward(args []interface{}) {
	m := args[0].(*message.Gate_Forward)
	//a := args[1].(gate.Agent)

	// 1. 获取玩家数据
	userData, ok := Users[m.UserID]
	if !ok {
		log.Error("玩家未在线", log.Int("UserID", m.UserID))
		return
	}

	// 2. 反序列化处理消息
	unmarshal, err := msg.JSONProcessor.Unmarshal(m.Body)
	if err != nil {
		log.Error("Unmarshal error", log.FieldErr(err), log.Int("UserID", m.UserID),
			log.String("Body", string(m.Body)))
		return
	}

	// 3. 路由
	err = msg.JSONProcessor.Route(unmarshal, userData)
	if err != nil {
		log.Error("route error", log.FieldErr(err), log.Int("UserID", m.UserID),
			log.String("Body", string(m.Body)))
	}
}

func Greeting(args []interface{}) {
	m := args[0].(*message.Greeting)
	ud := args[1].(*UserData)

	log.Info("Greeting ", m.Code, m.Message)

	m.Message = "welcome"
	b, _ := msg.JSONProcessor.Marshal(m)
	ud.Agent.WriteMsg(&message.S2C_Msg{
		From:  "",
		To:    "",
		MsgID: "",
		Agent: ud.AgentID,
		Body:  b[0],
	})

	if rand.Intn(100) > 80 {
		msg.JSONProcessor.Route(&message.Kick{
			UserID: ud.UserID,
			//Agent:  0,
		}, ud)
	}
}

func Disconnect(args []interface{}) {
	m := args[0].(*message.Disconnect)
	//a := args[1].(gate.Agent)

	_, ok := Users[m.UserID]
	if !ok {
		log.Error("玩家不在线，重复Disconnect", log.Int("UserID", m.UserID))
		return
	}

	delete(Users, m.UserID)
	log.Info("玩家下线", log.Int("UserID", m.UserID))
}

func Kick(args []interface{}) {
	m := args[0].(*message.Kick)
	//a := args[1].(gate.Agent)

	a, ok := Users[m.UserID]
	if !ok {
		log.Error("玩家不在线", log.Int("UserID", m.UserID))
		return
	}

	if a.Agent == nil {
		return
	}
	m.Agent = a.AgentID
	a.Agent.WriteMsg(m)
	log.Info("玩家被踢", log.Int("UserID", m.UserID))
}

func Login(args []interface{}) {
	m := args[0].(*message.Login)
	a := args[1].(gate.Agent)
	log.Info("Login server : ", m.Server, " account : ", m.Account, " agent : ", m.Agent)

	userID := m.Account
	_, ok := Users[userID]
	if ok {
		log.Error("重复登录", log.Int("UserID", userID))
	}

	Users[userID] = &UserData{
		Agent:   a,
		UserID:  userID,
		AgentID: m.Agent,
	}
	a.WriteMsg(&message.Bind{
		Agent:  m.Agent,
		UserID: userID,
		Server: serverID,
	})
}
