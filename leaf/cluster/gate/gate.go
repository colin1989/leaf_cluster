package gate

import (
	"reflect"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/cluster/config"
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/cluster/session"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
)

var gate *Gate

type Gate struct {
	config.WSConfig
	server    *protos.Server
	processor network.Processor
	chanRPC   *chanrpc.Server
	worldAddr string

	worldSession *session.Session
	NodeClients  map[int32]*network.WSClient
	NodeSessions map[int32]*session.Session
}

//var _ cluster.ClusterRPC = (*Gate)(nil)

func New(server *protos.Server, chanRpc *chanrpc.Server) *Gate {
	gate = &Gate{
		WSConfig: config.DefaultWSConfig(server.Address),
		server:   server,
		//processor:    processor,
		chanRPC:      chanRpc,
		NodeClients:  map[int32]*network.WSClient{},
		NodeSessions: map[int32]*session.Session{},
	}

	chanRpc.Register("NewAgent", rpcNewAgent)
	chanRpc.Register("CloseAgent", rpcCloseAgent)
	chanRpc.Register("NewWorldFunc", NewWorldFunc)
	chanRpc.Register("NewNodeFunc", NewNodeFunc)

	chanRpc.Register(reflect.TypeOf(&protos.Response{}), OnResponse)
	chanRpc.Register(reflect.TypeOf(&protos.WatchResponse{}), OnWatch)
	chanRpc.Register(reflect.TypeOf(&protos.Register{}), OnRegister)
	chanRpc.Register(reflect.TypeOf(&protos.Offline{}), OnOffline)
	chanRpc.Register(reflect.TypeOf(&protos.Bind{}), OnBind)
	chanRpc.Register(reflect.TypeOf(&protos.Kick{}), OnKick)

	return gate
}

func (g *Gate) Server() *protos.Server {
	return g.server
}

func (g *Gate) Listen(closeSig chan struct{}) {
	g.connectTo(0, g.worldAddr, "NewWorldFunc")
	<-closeSig
}

func (g *Gate) connectTo(sid int32, addr string, newSessionFunc string) *network.WSClient {
	wsClient := new(network.WSClient)
	wsClient.Addr = "ws://" + addr
	wsClient.ConnNum = 1
	wsClient.AutoReconnect = true
	wsClient.PendingWriteNum = g.PendingWriteNum
	wsClient.MaxMsgLen = g.MaxMsgLen
	wsClient.NewAgent = func(conn *network.WSConn) network.Agent {
		s := session.NewSession(conn, g.processor, g.chanRPC)
		if g.chanRPC != nil {
			g.chanRPC.Go(newSessionFunc, s, sid)
		}
		return s
	}

	if wsClient != nil {
		wsClient.Start()
	}

	return wsClient
}

func (g *Gate) Destroy() {
	log.InfoW("网关下线", log.Int32("serverID", g.server.ID))
	g.worldSession.WriteMsg(&protos.Offline{Server: g.server})
}

func (g *Gate) SetWorldAdd(worldAddr string) {
	g.worldAddr = worldAddr
}

func (g *Gate) SetProcessor(processor network.Processor) {
	g.processor = processor // NewGateProcessor(processor)
}

func (g *Gate) GetProcessor() network.Processor {
	return g.processor
}
