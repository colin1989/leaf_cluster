package master

import (
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/cluster/session"
	"github.com/name5566/leaf/log"
)

func OnCloseSession(args []interface{}) {
	s := args[0].(*session.Session)
	server := s.Server()
	if server == nil {
		log.Warn("服务在未注册前掉线")
		return
	}
	if !master.Set.Remove(server) {
		return
	}

	log.ErrorW("服务异常停止", log.Int32("serverID", server.ID),
		log.Int("type", int(server.Typ)), log.String("address", server.Address))

	if server.CheckType(protos.ServerType_Node) {
		master.nodeOffline(server)
	}
}

func OnRegister(args []interface{}) {
	m := args[0].(*protos.Register)
	s := args[1].(*session.Session)
	s.SetServer(m.Server)

	master.Set.AddServer(m.Server, s)

	log.InfoW("add server ", log.Int32("serverID", m.Server.ID),
		log.Int("type", int(m.Server.Typ)), log.String("address", m.Server.Address))

	if m.Server.GetTyp() == protos.ServerType_Gate {
		master.gateOnline(s)
	} else {
		master.nodeOnline(m.Server)
	}
}

func OnOffline(args []interface{}) {
	offline := args[0].(*protos.Offline)
	//s := args[1].(*session.Session)

	server := offline.Server

	if !master.Set.Remove(server) {
		log.InfoW("服务下线异常，不在列表中", log.Int32("serverID", server.ID),
			log.Int("type", int(server.Typ)), log.String("address", server.Address))
		return
	}

	log.InfoW("收到服务下线", log.Int32("serverID", server.ID),
		log.Int("type", int(server.Typ)), log.String("address", server.Address))

	if server.CheckType(protos.ServerType_Node) {
		master.nodeOffline(server)
	}
}

// gateOnline
//
//	@Description: 网关服上线，推送所有游戏服数据
//	@param gateAgent
func (m *Master) gateOnline(session *session.Session) {
	nodes := m.Set.AllNodes()
	if len(nodes) == 0 {
		return
	}
	session.WriteMsg(&protos.WatchResponse{
		Servers: nodes,
	})
}

// nodeOnline
//
//	@Description: 游戏服上线，通知所有网关服
//	@param game
func (m *Master) nodeOnline(node *protos.Server) {
	gates := m.Set.Gates()
	for _, sa := range gates {
		sa.Session.WriteMsg(&protos.WatchResponse{
			Servers: []*protos.Server{node},
		})
	}
}

func (m *Master) nodeOffline(node *protos.Server) {
	gates := m.Set.Gates()
	for _, sa := range gates {
		sa.Session.WriteMsg(&protos.Offline{
			Server: node,
		})
	}
}
