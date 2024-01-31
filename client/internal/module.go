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

	connectWorld()
}

func (m *Module) OnDestroy() {
	fmt.Println("module destroy")
}

func connectWorld() {
	client := Gate.Connect(1, "127.0.0.1:12345", "NewWorldServer", 1)
	client.AutoReconnect = false
}
