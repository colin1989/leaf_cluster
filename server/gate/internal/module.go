package internal

import (
	"fmt"
	"github.com/name5566/leaf/module"
	"server/base"
	"server/msg"
	"time"

	"github.com/name5566/leaf/gate"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

var (
	Gate = &gate.Gate{
		MaxConnNum:      20000,
		PendingWriteNum: 20000,
		MaxMsgLen:       40960,
		WSAddr:          "ws://127.0.0.1:13564",
		HTTPTimeout:     10 * time.Second,
		TCPAddr:         "",
		LenMsgLen:       2,
		LittleEndian:    false,
		Processor:       NewForwardProcessor(msg.JSONProcessor),
		AgentChanRPC:    ChanRPC,
	}
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	fmt.Println("gate module init")

	m.Skeleton = skeleton

	//连接Game服
	Gate.InitClient(1)
	//clients := Gate.InitClients(mapAddrs)
}

func (m *Module) OnDestroy() {
	fmt.Println("module destroy")
}
