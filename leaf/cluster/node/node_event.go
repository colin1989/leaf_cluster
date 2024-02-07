package node

import (
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/cluster/session"
	"github.com/name5566/leaf/log"
)

func NewWorldFunc(args []interface{}) {
	s := args[0].(*session.Session)

	s.WriteMsg(&protos.Register{Server: node.server})

	node.worldSession = s
}

func OnRequest(args []interface{}) {
	request := args[0].(*protos.Request)
	s := args[1].(*session.Session)

	sessionData := request.Session
	msg := request.Msg
	log.Infof("Request : %v", request)

	for _, data := range msg.Data {
		// 2. 反序列化处理消息
		unmarshal, err := node.processor.Unmarshal(data)
		if err != nil {
			log.ErrorW("Unmarshal error", log.FieldErr(err), log.Int64("UserID", sessionData.UId),
				log.String("Body", string(data)))
			return
		}

		// 3. 路由
		err = node.processor.Route(unmarshal, s, sessionData)
		if err != nil {
			log.ErrorW("route error", log.FieldErr(err), log.Int64("UserID", sessionData.UId),
				log.String("Body", string(data)))
		}
	}
}

func OnRegister(args []interface{}) {
	reg := args[0].(*protos.Register)
	s := args[1].(*session.Session)
	s.SetServer(reg.Server)
}
