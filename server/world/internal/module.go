package internal

import (
	"fmt"
	"server/base"

	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/cluster/protos"
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

	server := &protos.Server{
		ID:      1,
		Address: "127.0.0.1:12345",
		Typ:     protos.ServerType_Master,
	}
	cluster.NewMaster(server, ChanRPC)
}

func (m *Module) OnDestroy() {
	fmt.Println("world module destroy")
}
