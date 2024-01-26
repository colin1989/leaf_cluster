package internal

import (
	"client/msg"
	"time"

	"github.com/name5566/leaf/gate"
)

var (
	// 连接到服务器
	Gate = &gate.Gate{
		MaxConnNum:      20000,
		PendingWriteNum: 20000,
		MaxMsgLen:       40960,
		WSAddr:          "ws://127.0.0.1:13563",
		HTTPTimeout:     10 * time.Second,
		TCPAddr:         "",
		LenMsgLen:       2,
		LittleEndian:    false,
		Processor:       msg.JSONProcessor,
		AgentChanRPC:    ChanRPC,
	}

	// 接受其他节点的连接
	// Agent gate.Agent
)
