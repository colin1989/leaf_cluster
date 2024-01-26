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

var Client gate.Agent

func init() {
	// rpc
	handleMsg(&msg.S2S_Msg{}, S2S_Msg)

	// client msg
	handleMsg(&msg.Greeting{}, Greeting)
}

func S2S_Msg(args []interface{}) {
	m := args[0].(*msg.S2S_Msg)
	a := args[1].(gate.Agent) // game server
	if a == nil || m == nil {
		fmt.Println("greeting err")
		return
	}

	b := m.Body
	var greeting msg.Greeting
	_ = json.Unmarshal(b, &greeting)

	fmt.Println("from", m.From, "to", m.To, greeting.Code, greeting.Message)

	// 转发给客户端
	if Client != nil {
		fmt.Println("send to client")
		Client.WriteMsg(&greeting)
	}
}

func Greeting(args []interface{}) {
	m := args[0].(*msg.Greeting)
	a := args[1].(gate.Agent) // client
	_ = a

	if Client == nil {
		//TODO 将client agent通过session管理起来
		Client = a
	}

	fmt.Println("greeting", m.Code, m.Message)

	// RPC to game server
	if GameSvr != nil {
		fmt.Println("send to game server")
		b, _ := json.Marshal(&m)
		GameSvr.WriteMsg(&msg.S2S_Msg{
			From:  "gate",
			To:    "game",
			MsgID: msg.GetMsgId(&msg.Greeting{}),
			Body:  b,
		})
	}
}
