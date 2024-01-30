package internal

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"message"
	"reflect"

	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&message.Login{}, Login)
	handleMsg(&message.C2S_Msg{}, C2S_Msg)
	handleMsg(&message.Greeting{}, Greeting)
}

func C2S_Msg(args []interface{}) {
	m := args[0].(*message.C2S_Msg)
	a := args[1].(gate.Agent)
	_ = a

	log.Info("C2S_Msg ", m)
}

func Greeting(args []interface{}) {
	m := args[0].(*message.Greeting)
	a := args[1].(gate.Agent)
	_ = a

	log.Info("Greeting ", m.Code, m.Message)

	m.Message = "welcome"
	b, _ := json.Marshal(&m)
	_ = b
	a.WriteMsg(m)
}

func Login(args []interface{}) {
	m := args[0].(*message.Login)
	a := args[1].(gate.Agent)

	log.Info("Login server : ", m.Server, " account : ", m.Account, " agent : ", m.Agent)
	a.WriteMsg(&message.Bind{
		Agent:  m.Agent,
		UserID: m.Agent,
		Server: serverID,
	})
}
