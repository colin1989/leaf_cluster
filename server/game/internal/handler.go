package internal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.S2S_Msg{}, S2S_Msg)
}

func S2S_Msg(args []interface{}) {
	m := args[0].(*msg.S2S_Msg)
	a := args[1].(gate.Agent)

	var greeting msg.Greeting
	json.Unmarshal(m.Body, &greeting)

	fmt.Println("from", m.From, "to", m.To, m.MsgID, greeting.Code, greeting.Message)

	greeting.Message = "welcome"
	b, _ := json.Marshal(&greeting)
	a.WriteMsg(&msg.S2S_Msg{
		From: "game",
		To:   "gate",
		Body: b,
	})
}
