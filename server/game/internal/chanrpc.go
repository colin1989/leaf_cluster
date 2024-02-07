package internal

import (
	"fmt"

	"github.com/name5566/leaf/cluster/session"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseSession", rpcCloseAgent)
}

var Users = map[int64]*UserData{}

func rpcNewAgent(args []interface{}) {
	a := args[0].(*session.Session)
	_ = a
	//a.WriteMsg()

	fmt.Println("rpcNewAgent!!!")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(*session.Session)
	_ = a

	a.Close()

	fmt.Println("rpcCloseAgent!!!")
}
