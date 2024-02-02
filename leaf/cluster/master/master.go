package master

import (
	"reflect"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/cluster/config"
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/cluster/session"
	"github.com/name5566/leaf/network"
)

var master *Master

type Master struct {
	config.WSConfig
	server    *protos.Server
	processor network.Processor
	chanRPC   *chanrpc.Server

	Set *ServerSet
}

func New(server *protos.Server, chanRpc *chanrpc.Server) *Master {
	master = &Master{
		WSConfig: config.DefaultWSConfig(server.Address),
		server:   server,
		//processor: processor,
		chanRPC: chanRpc,
		Set:     NewServerSet(),
	}

	chanRpc.Register(reflect.TypeOf(&protos.Register{}), OnRegister)

	return master
}

func (m *Master) Server() *protos.Server {
	return m.server
}

func (m *Master) Listen(closeSig chan struct{}) {
	wsServer := new(network.WSServer)
	wsServer.Addr = m.WSAddr
	wsServer.MaxConnNum = m.MaxConnNum
	wsServer.PendingWriteNum = m.PendingWriteNum
	wsServer.MaxMsgLen = m.MaxMsgLen
	wsServer.HTTPTimeout = m.HTTPTimeout
	wsServer.NewAgent = func(conn *network.WSConn) network.Agent {
		s := session.NewSession(conn, m.processor, m.chanRPC)
		if m.chanRPC != nil {
			m.chanRPC.Go("NewSessionFunc", s)
		}
		return s
	}

	wsServer.Start()

	<-closeSig

	wsServer.Close()
}

func (m *Master) SetWorldAdd(worldAddr string) {
	//TODO implement me
	panic("implement me")
}

func (m *Master) SetProcessor(processor network.Processor) {
	m.processor = processor
}

func (m *Master) GetProcessor() network.Processor {
	return m.processor
}
