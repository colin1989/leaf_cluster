package master

import (
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/cluster/session"
	"github.com/name5566/leaf/log"
)

func OnRegister(args []interface{}) {
	m := args[0].(*protos.Register)
	s := args[1].(*session.Session)
	s.SetServer(m.Server)

	master.Set.AddServer(m.Server, s)

	log.Info("add server ", log.Int32("SId", m.Server.ID),
		log.Int("Type", int(m.Server.Typ)), log.String("Address", m.Server.Address))

	if m.Server.GetTyp() == protos.ServerType_Gate {
		master.gateOnline(s)
	} else {
		master.nodeOnline(m.Server)
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
