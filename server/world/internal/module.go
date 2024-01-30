package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"server/base"
	"server/world/msg"
	"time"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	fmt.Println("world module init")
	m.Gate = &gate.Gate{
		MaxConnNum:      20000,
		PendingWriteNum: 20000,
		MaxMsgLen:       40960,
		WSAddr:          "127.0.0.1:13560",
		HTTPTimeout:     10 * time.Second,
		TCPAddr:         "",
		LenMsgLen:       2,
		LittleEndian:    false,
		Processor:       msg.JSONProcessor,
		AgentChanRPC:    ChanRPC,
	}
}

func (m *Module) OnDestroy() {
	fmt.Println("module destroy")
}
