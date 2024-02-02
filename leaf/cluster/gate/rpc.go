package gate

import (
	"fmt"

	"github.com/name5566/leaf/cluster/protos"
)

func (g *Gate) RPCTo(sid int32, request *protos.Request) error {
	s, ok := g.NodeSessions[sid]
	if !ok {
		return fmt.Errorf("server %v does not online", sid)
	}
	s.WriteMsg(request)
	return nil
}
