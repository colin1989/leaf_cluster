package internal

import (
	"message"
	"reflect"
	"server/gate/msg"

	"github.com/name5566/leaf/cluster/protos"

	"github.com/name5566/leaf/agent"

	"github.com/name5566/leaf/rpc"

	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	//handleMsg(&message.C2S_Msg{}, C2S_Msg)
	handleMsg(&message.Login{}, Login)
}

func Login(args []interface{}) {
	m := args[0].(*message.Login)
	a := args[1].(agent.Agent) // game server
	if a == nil || m == nil {
		log.Infof("greeting err")
		return
	}

	s := a.UserData().(*protos.SessionData)
	if s == nil {
		log.Infof("agentData == nil")
		a.Close()
		return
	}

	if s.SId != 0 {
		log.Infof("agent has already logged int ")
		a.Close()
		return
	}

	m.Agent = s.AgentId

	data, err := msg.JSONProcessor.Marshal(m)
	if err != nil || len(data) < 1 {
		log.Errorf("JSONProcessor.Marshal login", log.FieldErr(err))
		return
	}
	if err = rpc.RPCLogin(m.Server, a, data); err != nil {
		log.Infof("SendToGame error : ", s.SId)
		a.Close()
		return
	}

}
