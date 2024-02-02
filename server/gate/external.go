package gate

import (
	"server/gate/internal"

	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/cluster/protos"
)

var (
	GateModule = new(internal.GateModule)
	Module     = new(internal.Module)
	ChanRPC    = internal.ChanRPC
)

func SetServerID(id int32) {
	internal.SetServerID(id)
}

func SetWSAddr(addr string) {
	internal.SetWSAddr(addr)
}

func init() {
	server := &protos.Server{
		ID:      1,
		Address: "127.0.0.1:13001",
		Typ:     protos.ServerType_Gate,
	}
	cluster.NewCluster(server, ChanRPC)
}
