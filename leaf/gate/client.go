package gate

import (
	"strconv"

	"github.com/name5566/leaf/network"
)

func (gate *Gate) InitClient(ConnNum int) {

	var wsClient *network.WSClient
	if gate.WSAddr != "" {
		wsClient = new(network.WSClient)
		wsClient.Addr = gate.WSAddr
		wsClient.ConnNum = ConnNum
		wsClient.AutoReconnect = true
		wsClient.PendingWriteNum = gate.PendingWriteNum
		wsClient.MaxMsgLen = gate.MaxMsgLen
		wsClient.NewAgent = func(conn *network.WSConn) network.Agent {
			a := &agent{conn: conn, gate: gate}
			if gate.AgentChanRPC != nil {
				gate.AgentChanRPC.Go("NewAgent", a)
			}
			return a
		}
	}

	if wsClient != nil {
		wsClient.Start()
	}
}

func (gate *Gate) InitClients(ConnAddrs map[string]string) (clients map[int]*network.WSClient) {

	clients = map[int]*network.WSClient{}

	for i, addr := range ConnAddrs {
		id, _ := strconv.Atoi(i)

		var wsClient *network.WSClient

		wsClient = new(network.WSClient)
		wsClient.Addr = "ws://" + addr
		wsClient.ConnNum = 1
		wsClient.AutoReconnect = true
		wsClient.PendingWriteNum = gate.PendingWriteNum
		wsClient.MaxMsgLen = gate.MaxMsgLen
		wsClient.NewAgent = func(conn *network.WSConn) network.Agent {
			a := &agent{conn: conn, gate: gate}
			if gate.AgentChanRPC != nil {
				gate.AgentChanRPC.Go("NewAgent", a, id)
			}
			return a
		}

		wsClient.Start()
		clients[id] = wsClient
	}

	return
}
