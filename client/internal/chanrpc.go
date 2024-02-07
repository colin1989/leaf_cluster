package internal

import (
	"fmt"
	"math/rand"
	"message"
	"reflect"
	"time"

	"github.com/name5566/leaf/agent"
	"github.com/name5566/leaf/log"
)

func init() {
	//skeleton.RegisterChanRPC("NewWorldServer", NewWorldServer)
	skeleton.RegisterChanRPC("NewGameServer", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)

	skeleton.RegisterChanRPC(reflect.TypeOf(&message.Login{}), Login)
	// skeleton.RegisterChanRPC(reflect.TypeOf(&message.S2C_Gates{}), S2C_Gates)
	skeleton.RegisterChanRPC(reflect.TypeOf(&message.Greeting{}), Greeting)
}

//func NewWorldServer(args []interface{}) {
//	a := args[0].(gate.Agent)
//
//	log.Info("Connect to world")
//
//	//a.WriteMsg(&message.C2S_Gates{})
//	fmt.Println("NewWorldServer!!!")
//}

var Account = rand.Intn(100)
var Code = rand.Intn(10000)

func rpcNewAgent(args []interface{}) {
	a := args[0].(agent.Agent)

	log.Info("login", log.Int("Account", Account))

	skeleton.AfterFunc(time.Second, func() {
		a.WriteMsg(&message.Login{
			Server:  int32(ServerID),
			Account: Account,
		})
	})
	fmt.Println("rpcNewAgent!!!")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(agent.Agent)
	_ = a

	fmt.Println("rpcCloseAgent!!!")
}

func Login(args []interface{}) {
	m := args[0].(*message.Login)
	a := args[1].(agent.Agent)

	if a == nil || m == nil {
		fmt.Println("Login")
		return
	}

	skeleton.Go(func() {
		time.Sleep(time.Second)
		a.WriteMsg(&message.Greeting{
			Code:    Code + 1,
			Message: "hello from client",
		})
	}, func() {
		log.Infof("Account %v Login Server %v", Account, ServerID)
	})
}

func Greeting(args []interface{}) {
	m := args[0].(*message.Greeting)
	a := args[1].(agent.Agent)

	if a == nil || m == nil {
		fmt.Println("greeting err")
		return
	}

	log.Infof("Account %v Server %v received Code %v Message %v", Account, ServerID, m.Code, m.Message)

	if Code != m.Code-1 {
		log.FatalW("Code != m.Code-1", log.Int("account", Account), log.Int("serverID", ServerID),
			log.Int("Code", m.Code), log.String("Message", m.Message))
	}
	Code = m.Code

	time.Sleep(time.Second)
	a.WriteMsg(&message.Greeting{
		Code:    Code + 1,
		Message: "hello from client",
	})

	log.Infof("Account %v Server %v Send Code %v", Account, ServerID, m.Code+1)

}
