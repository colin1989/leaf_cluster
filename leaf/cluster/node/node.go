package node

import (
	"reflect"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/cluster/config"
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/cluster/session"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
)

type Node struct {
	config.WSConfig
	server       *protos.Server
	processor    network.Processor
	chanRPC      *chanrpc.Server
	worldAddr    string
	worldSession *session.Session
}

var node *Node = nil

func New(server *protos.Server, chanRpc *chanrpc.Server) *Node {
	node = &Node{
		WSConfig: config.DefaultWSConfig(server.Address),
		server:   server,
		//processor: processor,
		chanRPC: chanRpc,
	}

	//chanRpc.Register("NewSessionFunc", NewSessionFunc)
	chanRpc.Register("NewWorldFunc", NewWorldFunc)

	chanRpc.Register(reflect.TypeOf(&protos.Request{}), OnRequest)
	chanRpc.Register(reflect.TypeOf(&protos.Register{}), OnRegister)
	return node
}

func (n *Node) Server() *protos.Server {
	return n.server
}

func (n *Node) Listen(closeSig chan struct{}) {
	wsServer := new(network.WSServer)
	wsServer.Addr = n.WSAddr
	wsServer.MaxConnNum = n.MaxConnNum
	wsServer.PendingWriteNum = n.PendingWriteNum
	wsServer.MaxMsgLen = n.MaxMsgLen
	wsServer.HTTPTimeout = n.HTTPTimeout
	wsServer.NewAgent = func(conn *network.WSConn) network.Agent {
		s := session.NewSession(conn, n.processor, n.chanRPC)
		if n.chanRPC != nil {
			n.chanRPC.Go("NewSessionFunc", s)
		}
		s.SetUserData("gate")
		return s
	}

	wsServer.Start()

	n.connectToWorld(n.worldAddr)
	<-closeSig

	wsServer.Close()
}

func (n *Node) connectToWorld(worldAddr string) *network.WSClient {
	wsClient := new(network.WSClient)
	wsClient.Addr = "ws://" + worldAddr
	wsClient.ConnNum = 1
	wsClient.AutoReconnect = true
	wsClient.PendingWriteNum = n.PendingWriteNum
	wsClient.MaxMsgLen = n.MaxMsgLen
	wsClient.NewAgent = func(conn *network.WSConn) network.Agent {
		s := session.NewSession(conn, n.processor, n.chanRPC)
		if n.chanRPC != nil {
			n.chanRPC.Go("NewWorldFunc", s)
		}
		return s
	}

	if wsClient != nil {
		wsClient.Start()
	}

	return wsClient
}

func (n *Node) Destroy() {
	log.InfoW("node 下线", log.Int32("serverID", n.server.ID))
	n.worldSession.WriteMsg(&protos.Offline{Server: n.server})
}

func (n *Node) SetWorldAdd(worldAddr string) {
	n.worldAddr = worldAddr
}

func (n *Node) SetProcessor(processor network.Processor) {
	n.processor = processor
}

func (n *Node) GetProcessor() network.Processor {
	return n.processor
}
