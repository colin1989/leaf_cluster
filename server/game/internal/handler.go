package internal

import (
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
	handleMsg(&message.Login{}, Login)
	handleMsg(&message.Greeting{}, Greeting)
}

func Greeting(args []interface{}) {
	m := args[0].(*message.Greeting)
	a := args[1].(gate.Agent)

	userData, ok := a.UserData().(*UserData)
	if !ok {
		log.Error("user has not logged")
		return
	}

	log.Info("Greeting ", m.Code, m.Message)

	m.Message = "welcome"
	b, _ := msg.JSONProcessor.Marshal(m)
	a.WriteMsg(&message.S2C_Msg{
		From:  "",
		To:    "",
		MsgID: "",
		Agent: userData.AgentID,
		Body:  b[0],
	})
}

func Login(args []interface{}) {
	m := args[0].(*message.Login)
	a := args[1].(gate.Agent)
	a.SetUserData(&UserData{
		UserID:  m.Agent,
		AgentID: m.Agent,
	})
	log.Info("Login server : ", m.Server, " account : ", m.Account, " agent : ", m.Agent)
	a.WriteMsg(&message.Bind{
		Agent:  m.Agent,
		UserID: m.Agent,
		Server: serverID,
	})
}
