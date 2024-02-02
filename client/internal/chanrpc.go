package internal

import (
	"fmt"
	"message"
	"reflect"
	"time"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	//skeleton.RegisterChanRPC("NewWorldServer", NewWorldServer)
	skeleton.RegisterChanRPC("NewGameServer", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)

	skeleton.RegisterChanRPC(reflect.TypeOf(&message.Login{}), Login)
	//skeleton.RegisterChanRPC(reflect.TypeOf(&message.S2C_Gates{}), S2C_Gates)
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

var Account = 1

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	serverID := args[1].(int32)

	log.Info("login", log.Int("Account", Account))

	skeleton.AfterFunc(time.Second, func() {
		a.WriteMsg(&message.Login{
			Server:  serverID,
			Account: Account,
		})
	})
	fmt.Println("rpcNewAgent!!!")

	Account++
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	connectWorld()
	fmt.Println("rpcCloseAgent!!!")
}

//func S2C_Gates(args []interface{}) {
//	m := args[0].(*message.S2C_Gates)
//	//a := args[1].(gate.Agent)
//
//	gameID := int32(0)
//	addr := ""
//	rand.Shuffle(len(m.Addresses), func(i, j int) {
//		tmp := m.Addresses[i]
//		m.Addresses[i] = m.Addresses[j]
//		m.Addresses[j] = tmp
//	})
//	rand.Shuffle(len(m.GameID), func(i, j int) {
//		tmp := m.GameID[i]
//		m.GameID[i] = m.GameID[j]
//		m.GameID[j] = tmp
//	})
//	if len(m.GameID) > 0 {
//		gameID = m.GameID[0]
//	}
//	if len(m.Addresses) > 0 {
//		addr = m.Addresses[0]
//	}
//
//	if addr == "" || gameID == 0 {
//		return
//	}
//	log.Info("Connect to server", log.String("gate", addr),
//		log.Int32("ServerID", gameID), log.Int("Account", Account))
//	client := Gate.Connect(gameID, addr, "NewGameServer", 1)
//	client.AutoReconnect = false
//}

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

	//if m.Code > 10 {
	//	return
	//}

	time.Sleep(time.Second)
	a.WriteMsg(&message.Greeting{
		Code:    m.Code + 1,
		Message: "hello from client",
	})

}
