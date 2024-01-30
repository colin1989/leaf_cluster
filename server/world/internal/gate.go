package internal

import (
	"fmt"
	"server/world/msg"
	"time"

	"github.com/name5566/leaf/gate"
)

type GateModule struct {
	*gate.Gate
}

func (m *GateModule) OnInit() {
	fmt.Println("world gate module init")

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

func (m *GateModule) OnDestroy() {
	fmt.Println("world gate module destroy")
}
