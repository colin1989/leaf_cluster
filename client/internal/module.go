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

	//连接Gate服
	Gate.InitClient(1)
}

func (m *Module) OnDestroy() {
	fmt.Println("module destroy")
}
