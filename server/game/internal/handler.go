package internal

import (
	"math/rand"
	"message"
	"reflect"

	"github.com/name5566/leaf/cluster/protos"

	"github.com/name5566/leaf/cluster/session"

	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&protos.Disconnect{}, Disconnect)
	handleMsg(&message.Login{}, Login)
	handleMsg(&message.Greeting{}, Greeting)
}

func Greeting(args []interface{}) {
	m := args[0].(*message.Greeting)
	s := args[1].(*session.Session)
	sd := args[2].(*protos.SessionData)

	log.Info("Greeting ", m.Code, m.Message)

	m.Message = "welcome"
	//b, _ := msg.JSONProcessor.Marshal(m)
	s.WriteResponse(m, sd)

	//sd := s.SessionData()

	if rand.Intn(50) > 80 {
		s.Kick(sd.UId, sd.AgentId)
		//msg.JSONProcessor.Route(&message.Kick{
		//	UserID: sd.UId,
		//	//Agent:  0,
		//}, sd)
	}
}

func Disconnect(args []interface{}) {
	m := args[0].(*protos.Disconnect)
	//a := args[1].(gate.Agent)

	_, ok := Users[m.UId]
	if !ok {
		log.Error("玩家不在线，重复Disconnect", log.Int64("UserID", m.UId))
		return
	}

	delete(Users, m.UId)
	log.Info("玩家下线", log.Int64("UserID", m.UId))
}

func Login(args []interface{}) {
	m := args[0].(*message.Login)
	s := args[1].(*session.Session)
	sd := args[2].(*protos.SessionData)
	log.Info("Login server : ", m.Server, " account : ", m.Account, " agent : ", m.Agent)

	userID := int64(m.Account)
	_, ok := Users[userID]
	if ok {
		log.Error("重复登录", log.Int64("UserID", userID))
	}

	Users[userID] = &UserData{
		Agent:   s,
		UserID:  userID,
		AgentID: m.Agent,
	}

	sd.AgentId = m.Agent
	sd.UId = userID
	if err := s.Bind(m.Agent, userID, serverID); err != nil {
		log.Errorf("玩家绑定数据失败", log.Int64("agent", m.Agent),
			log.Int64("userID", userID))
	}

	s.WriteResponse(m, sd)
}
