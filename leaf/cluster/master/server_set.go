package master

import (
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/cluster/session"
	"github.com/name5566/leaf/log"
)

type ServerAgent struct {
	Server  *protos.Server
	Session *session.Session
}

type ServerSet struct {
	gates []*ServerAgent
	nodes map[int32]*ServerAgent
}

func NewServerSet() *ServerSet {
	ss := &ServerSet{
		gates: []*ServerAgent{},
		nodes: map[int32]*ServerAgent{},
	}

	return ss
}

func (ss *ServerSet) AddServer(server *protos.Server, s *session.Session) {
	sa := &ServerAgent{
		Server:  server,
		Session: s,
	}
	switch server.Typ {
	case protos.ServerType_Gate:
		for i, gate := range ss.gates {
			if gate.Server.ID != server.ID {
				continue
			}
			ss.gates[i] = sa
			log.InfoW("覆盖添加网关节点", log.Int32("ServerID", server.ID))
			return
		}
		ss.gates = append(ss.gates, sa)
	case protos.ServerType_Node:
		_, ok := ss.nodes[server.ID]
		if ok {
			log.InfoW("覆盖Node节点", log.Int32("ServerID", server.ID))
		}
		ss.nodes[server.ID] = sa
		log.InfoW("游戏服节点添加", log.Int32("ServerID", server.ID))
	case protos.ServerType_Master:
		log.Error("Master add master server")
	}
}

func (ss *ServerSet) Gates() []*ServerAgent {
	//gates := make([]*protos.Server, len(ss.nodes))
	//for _, gate := range ss.gates {
	//	gates = append(gates, gate.Server)
	//}
	return ss.gates
}

func (ss *ServerSet) Node(id int32) (*ServerAgent, bool) {
	s, ok := ss.nodes[id]
	return s, ok
}

func (ss *ServerSet) AllNodes() []*protos.Server {
	games := make([]*protos.Server, 0, len(ss.nodes))
	for _, serverAgent := range ss.nodes {
		games = append(games, serverAgent.Server)
	}
	return games
}

func (ss *ServerSet) Remove(s *protos.Server) bool {
	switch s.Typ {
	case protos.ServerType_Gate:
		for i, gate := range ss.gates {
			if gate.Server.ID != s.ID {
				continue
			}
			ss.gates = append(ss.gates[:i], ss.gates[i+1:]...)
			return true
		}
	case protos.ServerType_Node:
		_, ok := ss.nodes[s.ID]
		if !ok {
			return false
		}
		delete(ss.nodes, s.ID)
		return true
	}
	return false
}
