package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"message"
	"reflect"
	"server/protos"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&message.S2S_Reg{}, S2S_Reg)
}

func S2S_Reg(args []interface{}) {
	m := args[0].(*message.S2S_Reg)
	a := args[1].(gate.Agent)
	a.SetUserData(m.Server)

	manager.AddServer(m.Server, a)

	log.Info("add server ", log.Int32("SId", m.Server.ID),
		log.Int("Type", int(m.Server.Typ)), log.String("Address", m.Server.Address))

	if m.Server.GetTyp() == protos.ServerType_Gate {
		gateOnline(a)
	} else {
		gameOnline(m.Server)
	}
}

func gameOnline(game *protos.Server) {
	gates := manager.GetServers(protos.ServerType_Gate)
	for _, sa := range gates {
		sa.Agent.WriteMsg(&message.W2S_GS{
			Servers: []*protos.Server{game},
		})
	}
}

func gateOnline(gateAgent gate.Agent) {
	game := manager.GetServers(protos.ServerType_Game)
	servers := make([]*protos.Server, 0, len(game))
	for _, sa := range game {
		servers = append(servers, sa.Server)
	}
	gateAgent.WriteMsg(&message.W2S_GS{
		Servers: servers,
	})
}