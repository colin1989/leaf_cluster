package cluster

import (
	"log"
	"reflect"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/cluster/gate"
	"github.com/name5566/leaf/cluster/master"
	"github.com/name5566/leaf/cluster/node"
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
	"github.com/name5566/leaf/network/protobuf"
)

type Cluster struct {
	Node ClusterNode

	stop chan struct{}
}

func NewMaster(server *protos.Server, chanRpc *chanrpc.Server) *Cluster {
	if defaultCluster != nil {
		panic("重复创建Cluster")
	}
	if server.Typ != protos.ServerType_Master {
		panic("sever type error")
	}

	clusterNode := master.New(server, chanRpc)
	clusterNode.SetProcessor(json.NewProcessor())

	listenMessage(clusterNode.GetProcessor(), chanRpc)

	defaultCluster = &Cluster{
		Node: clusterNode,
		stop: make(chan struct{}),
	}

	return defaultCluster
}

// NewGate
//
//	@Description: 创建网关
//	@param server
//	@param processor
//	@param chanRpc
func NewGate(server *protos.Server, chanRpc *chanrpc.Server, worldAddr string) *Cluster {
	if defaultCluster != nil {
		panic("重复创建Cluster")
	}

	g := gate.New(server, chanRpc)

	// defaultProcessor
	g.SetProcessor(json.NewProcessor())
	g.SetWorldAdd(worldAddr)
	listenMessage(g.GetProcessor(), chanRpc)

	defaultCluster = &Cluster{
		Node: g,
		stop: make(chan struct{}),
	}

	return defaultCluster
}

// NewNode
//
//	@Description: 根据传入的服务类型，创建对应的节点
//	@param server
//	@param processor
//	@param chanRpc
func NewNode(server *protos.Server, chanRpc *chanrpc.Server, worldAdd string, processor network.Processor) *Cluster {
	if defaultCluster != nil {
		panic("重复创建Cluster")
	}
	n := node.New(server, chanRpc)

	// defaultProcessor
	n.SetWorldAdd(worldAdd)
	n.SetProcessor(processor)

	listenMessage(n.GetProcessor(), chanRpc)

	defaultCluster = &Cluster{
		Node: n,
		stop: make(chan struct{}),
	}

	return defaultCluster
}

func listenMessage(processor network.Processor, chanRpc *chanrpc.Server) {
	switch p := processor.(type) {
	case *protobuf.Processor:
		listenProto(p, chanRpc)
	case *json.Processor:
		listenJSON(p, chanRpc)
	default:
		msgType := reflect.TypeOf(p)
		log.Panicf("processor type error %v", msgType.Elem().Name())
	}
}

func listenProto(processor *protobuf.Processor, chanRpc *chanrpc.Server) {
	processor.Register(&protos.Register{})
	processor.SetRouter(&protos.Register{}, chanRpc)

	processor.Register(&protos.Request{})
	processor.SetRouter(&protos.Request{}, chanRpc)

	processor.Register(&protos.Response{})
	processor.SetRouter(&protos.Response{}, chanRpc)

	processor.Register(&protos.WatchResponse{})
	processor.SetRouter(&protos.WatchResponse{}, chanRpc)

	processor.Register(&protos.Bind{})
	processor.SetRouter(&protos.Bind{}, chanRpc)

	processor.Register(&protos.Kick{})
	processor.SetRouter(&protos.Kick{}, chanRpc)

	processor.Register(&protos.Disconnect{})
	processor.SetRouter(&protos.Disconnect{}, chanRpc)
}

func listenJSON(processor *json.Processor, chanRpc *chanrpc.Server) {
	processor.Register(&protos.Register{})
	processor.SetRouter(&protos.Register{}, chanRpc)

	processor.Register(&protos.Request{})
	processor.SetRouter(&protos.Request{}, chanRpc)

	processor.Register(&protos.Response{})
	processor.SetRouter(&protos.Response{}, chanRpc)

	processor.Register(&protos.WatchResponse{})
	processor.SetRouter(&protos.WatchResponse{}, chanRpc)

	processor.Register(&protos.Bind{})
	processor.SetRouter(&protos.Bind{}, chanRpc)

	processor.Register(&protos.Kick{})
	processor.SetRouter(&protos.Kick{}, chanRpc)

	processor.Register(&protos.Disconnect{})
	processor.SetRouter(&protos.Disconnect{}, chanRpc)
}

func GetNode() ClusterNode {
	return defaultCluster.Node
}

func (c *Cluster) listen() {
	c.Node.Listen(defaultCluster.stop)
}
