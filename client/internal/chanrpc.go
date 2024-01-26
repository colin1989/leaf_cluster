package internal

import (
	"client/msg"
	"fmt"
	"reflect"
	"time"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)

	skeleton.RegisterChanRPC(reflect.TypeOf(&msg.Greeting{}), Greeting)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	skeleton.Go(func() {
		time.Sleep(time.Second)
		a.WriteMsg(&msg.Greeting{
			Code:    1,
			Message: "hello from client",
		})
	}, func() {
		fmt.Println("write msg")
	})

	fmt.Println("rpcNewAgent!!!")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	fmt.Println("rpcCloseAgent!!!")
}

func Greeting(args []interface{}) {
	m := args[0].(*msg.Greeting)
	a := args[1].(gate.Agent)

	if a == nil || m == nil {
		fmt.Println("greeting err")
		return
	}

	fmt.Println("received", m.Code, m.Message)

	if m.Code > 10 {
		return
	}

	time.Sleep(time.Second)
	a.WriteMsg(&msg.Greeting{
		Code:    m.Code + 1,
		Message: "hello from client",
	})

}
