package cluster

import (
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/network"
)

//goland:noinspection GoNameStartsWithPackageName
type ClusterNode interface {
	Server() *protos.Server
	Listen(closeSig chan struct{})
	GetProcessor() network.Processor
}

type ClusterRPC interface {
	RPCTo(sid int32, request *protos.Request) error
}
