package rpc

import (
	"fmt"

	"github.com/name5566/leaf/agent"

	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/cluster/protos"
)

//goland:noinspection GoNameStartsWithPackageName
func RPCLogin(sid int32, a agent.Agent, data [][]byte) error {
	if cluster.GetNode() == nil {
		return fmt.Errorf("cluster is nil")
	}

	rpcServer, ok := cluster.GetNode().(cluster.ClusterRPC)
	if !ok {
		return fmt.Errorf("node does not implement ClusterRPC")
	}

	s := &protos.SessionData{
		AgentId: a.ID(),
		UId:     0,
		Data:    nil,
	}

	m := &protos.Msg{
		Id:    0,
		Route: "",
		Data:  data,
		Type:  protos.MsgType_MsgLogin,
	}

	req := &protos.Request{
		Session: s,
		Msg:     m,
		Server:  cluster.GetNode().Server(),
	}

	return rpcServer.RPCTo(sid, req)
}
