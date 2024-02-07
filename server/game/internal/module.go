package internal

import (
	"fmt"
	"server/base"
	"server/game/msg"

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
	fmt.Println("game module init")

	m.Skeleton = skeleton

	server := &protos.Server{
		ID:      serverID,
		Address: wsAddr,
		Typ:     protos.ServerType_Node,
	}

	worldAddr := "127.0.0.1:12345"
	cluster.NewNode(server, ChanRPC, worldAddr, msg.JSONProcessor)
	//cluster.WithWorld(worldAddr),
	//cluster.WithProcessor(msg.JSONProcessor))
}

func (m *Module) OnDestroy() {
	fmt.Println("module destroy")
}
