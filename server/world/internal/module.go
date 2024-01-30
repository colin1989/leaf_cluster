package internal

import (
	"fmt"
	"server/base"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	fmt.Println("world module init")

	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {
	fmt.Println("world module destroy")
}
