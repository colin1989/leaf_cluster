package gate

import (
	"github.com/name5566/leaf/agent"
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/cluster/session"
	"github.com/name5566/leaf/log"
)

var AgentMap = map[int64]agent.Agent{}

func rpcNewAgent(args []interface{}) {
	a := args[0].(agent.Agent)
	agentId := args[1].(int64)
	a.SetUserData(&protos.SessionData{
		AgentId: agentId,
		UId:     0,
		SId:     0,
		Data:    nil,
	})
	AgentMap[agentId] = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(agent.Agent)

	delete(AgentMap, a.ID())

	sd, ok := a.UserData().(*protos.SessionData)
	if !ok {
		return
	}
	s, ok := gate.NodeSessions[sd.SId]
	if !ok {
		log.Errorf("server %v does not online", sd.SId)
		return
	}
	s.WriteMsg(&protos.Disconnect{
		AgentId: sd.AgentId,
		UId:     sd.UId,
	})
}

func NewWorldFunc(args []interface{}) {
	s := args[0].(*session.Session)

	s.WriteMsg(&protos.Register{Server: gate.server})

	gate.worldSession = s
	log.Info("gate connect world")
}

func NewNodeFunc(args []interface{}) {
	s := args[0].(*session.Session)
	sid := args[1].(int32)

	_, ok := gate.NodeSessions[sid]
	if ok {
		log.ErrorW("重复连接", log.Int32("sid", sid))
	}
	gate.NodeSessions[sid] = s
	log.InfoW("gate connect node", log.Int32("sid", sid))

	s.WriteMsg(&protos.Register{Server: gate.server})
}

func OnWatch(args []interface{}) {
	m := args[0].(*protos.WatchResponse)
	//s := args[1].(*session.Session)

	for _, server := range m.Servers {
		// TODO 这个地方可能会有问题
		if client, ok := gate.NodeClients[server.ID]; ok {
			// Close 里面有个wait...
			go client.Close()
		}
		if _, ok := gate.NodeSessions[server.ID]; ok {
			gate.NodeSessions[server.ID].Close()
			delete(gate.NodeSessions, server.ID)
		}
		gate.NodeClients[server.ID] = gate.connectTo(server.ID, server.Address, "NewNodeFunc")
	}
}

func OnResponse(args []interface{}) {
	response := args[0].(*protos.Response)
	//s := args[1].(*session.Session)

	a, ok := AgentMap[response.Session.AgentId]
	if !ok {
		log.Errorf("agent 数据丢失", response.Session.AgentId)
		return
	}

	for _, datum := range response.Msg.Data {
		a.WriteRaw(datum)
	}
}

func OnRegister(args []interface{}) {
	m := args[0].(*protos.Register)
	s := args[1].(*session.Session)

	s.SetServer(m.Server)
	log.InfoW("服务节点注册", log.Int32("serverID", m.Server.ID),
		log.Int32("type", int32(m.Server.Typ)))
}

func OnOffline(args []interface{}) {
	m := args[0].(*protos.Offline)
	//s := args[1].(*session.Session)

	server := m.Server
	if client, ok := gate.NodeClients[server.ID]; ok {
		// Close 里面有个wait...
		go client.Close()
	} else {
		log.WarnW("服务下线, NodeClients数据不存在", log.Int32("serverID", server.ID))
	}

	if _, ok := gate.NodeSessions[server.ID]; ok {
		gate.NodeSessions[server.ID].Close()
		delete(gate.NodeSessions, server.ID)
	} else {
		log.WarnW("服务下线, NodeSessions数据不存在", log.Int32("serverID", server.ID))
	}
	log.InfoW("服务节点下线", log.Int32("serverID", m.Server.ID),
		log.Int32("type", int32(m.Server.Typ)))
}

func OnBind(args []interface{}) {
	m := args[0].(*protos.Bind)
	//s := args[1].(*session.Session)

	a, ok := AgentMap[m.AgentId]
	if !ok {
		log.Errorf("agent 数据丢失", m.AgentId)
		return
	}

	sd, ok := a.UserData().(*protos.SessionData)
	if !ok {
		return
	}

	sd.UId = m.UId
	sd.SId = m.SId
	log.InfoW("用户数据Bind", log.Int64("agentID", m.AgentId),
		log.Int64("userID", sd.UId), log.Int32("serverID", sd.SId))
}

func OnKick(args []interface{}) {
	k := args[0].(*protos.Kick)
	//s := args[1].(*session.Session)

	a, ok := AgentMap[k.AgentId]
	if !ok {
		log.Errorf("agent 数据丢失", k.AgentId)
		return
	}

	a.Close()
	log.InfoW("Kick", log.Int64("agentID", k.AgentId))
}
