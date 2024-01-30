package internal

import (
	"server/protos"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type ServerAgent struct {
	Server *protos.Server
	Agent  gate.Agent
}

var manager = &ServerManager{
	Servers: map[protos.ServerType]map[int32]*ServerAgent{},
}

type ServerManager struct {
	Servers map[protos.ServerType]map[int32]*ServerAgent
}

func (sm *ServerManager) ensureServers(typ protos.ServerType) {
	_, ok := sm.Servers[typ]
	if !ok {
		sm.Servers[typ] = map[int32]*ServerAgent{}
	}
}

func (sm *ServerManager) AddServer(s *protos.Server, a gate.Agent) {
	sm.ensureServers(s.Typ)
	_, ok := sm.Servers[s.Typ][s.ID]
	if ok {
		log.Warn("add server, already exits", log.Int32("SId", s.ID))
	}
	sm.Servers[s.Typ][s.ID] = &ServerAgent{
		Server: s,
		Agent:  a,
	}
}

func (sm *ServerManager) RemoveServer(typ protos.ServerType, id int32) {
	sm.ensureServers(typ)
	delete(sm.Servers[typ], id)
}

func (sm *ServerManager) GetServers(typ protos.ServerType) map[int32]*ServerAgent {
	sm.ensureServers(typ)
	return sm.Servers[typ]
}
