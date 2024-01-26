package internal

import (
	"fmt"
	"github.com/name5566/leaf/module"
	"server/base"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	fmt.Println("game module init")

	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {
	fmt.Println("module destroy")
}
