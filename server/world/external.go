package world

import (
	"server/world/internal"

	"github.com/name5566/leaf/cluster/protos"

	"github.com/name5566/leaf/cluster"
)

var (
	Module = new(internal.Module)
	//GateModule = new(internal.GateModule)
	ChanRPC = internal.ChanRPC
)

func init() {
	server := &protos.Server{
		ID:      1,
		Address: "127.0.0.1:12345",
		Typ:     protos.ServerType_Master,
	}
	cluster.NewCluster(server, ChanRPC)
}
