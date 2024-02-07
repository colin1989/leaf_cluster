package internal

import (
	"fmt"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	fmt.Println("module init")

	m.Skeleton = skeleton

	addr := GateAddr
	if len(addr) == 0 {
		addr = "127.0.0.1:13561"
	}
	client := Gate.Connect(1, addr, "NewGameServer", 1)
	client.AutoReconnect = false
	//connectWorld()
}

func (m *Module) OnDestroy() {
	fmt.Println("module destroy")
}
