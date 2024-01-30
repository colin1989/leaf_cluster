package internal

import (
	"fmt"
	"message"
	"reflect"
	"time"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewGameServer", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)

	skeleton.RegisterChanRPC(reflect.TypeOf(&message.Login{}), Login)
	skeleton.RegisterChanRPC(reflect.TypeOf(&message.Greeting{}), Greeting)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	skeleton.AfterFunc(time.Second, func() {
		a.WriteMsg(&message.Login{
			Server:  1,
			Account: 1,
		})
	})
	fmt.Println("rpcNewAgent!!!")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	fmt.Println("rpcCloseAgent!!!")
}

func Login(args []interface{}) {
	m := args[0].(*message.Login)
	a := args[1].(gate.Agent)

	if a == nil || m == nil {
		fmt.Println("Login")
		return
	}

	skeleton.Go(func() {
		time.Sleep(time.Second)
		a.WriteMsg(&message.Greeting{
			Code:    1,
			Message: "hello from client",
		})
	}, func() {
		fmt.Println("Login")
	})

}

func Greeting(args []interface{}) {
	m := args[0].(*message.Greeting)
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
	a.WriteMsg(&message.Greeting{
		Code:    m.Code + 1,
		Message: "hello from client",
	})

}
