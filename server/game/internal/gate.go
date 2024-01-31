package internal

import (
	"fmt"
	"server/game/msg"
	"time"

	"github.com/name5566/leaf/gate"
)

type GateModule struct {
	*gate.Gate
}

func (m *GateModule) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      20000,
		PendingWriteNum: 20000,
		MaxMsgLen:       40960,
		WSAddr:          wsAddr,
		HTTPTimeout:     10 * time.Second,
		TCPAddr:         "",
		LenMsgLen:       2,
		LittleEndian:    false,
		//Processor:       NewGameMsgProcessor(msg.JSONProcessor),
		Processor:    msg.JSONProcessor,
		AgentChanRPC: ChanRPC,
	}
	m.Gate.Connect(0, "127.0.0.1:12345", NewWorldAgentFunc, 1)
}

func (m *GateModule) OnDestroy() {
	fmt.Println("module destroy")
}
