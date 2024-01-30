package gate

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/util/compress"
)

type Gate struct {
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	Processor       network.Processor
	AgentChanRPC    *chanrpc.Server

	// websocket
	WSAddr      string
	HTTPTimeout time.Duration

	// tcp
	TCPAddr      string
	LenMsgLen    int
	LittleEndian bool
}

func (gate *Gate) Run(closeSig chan bool) {
	var wsServer *network.WSServer
	if gate.WSAddr != "" {
		wsServer = new(network.WSServer)
		wsServer.Addr = gate.WSAddr
		wsServer.MaxConnNum = gate.MaxConnNum
		wsServer.PendingWriteNum = gate.PendingWriteNum
		wsServer.MaxMsgLen = gate.MaxMsgLen
		wsServer.HTTPTimeout = gate.HTTPTimeout
		wsServer.NewAgent = func(conn *network.WSConn) network.Agent {
			a := &agent{conn: conn, gate: gate}
			if gate.AgentChanRPC != nil {
				gate.AgentChanRPC.Go("NewAgent", a)
			}
			return a
		}
	}

	var tcpServer *network.TCPServer
	if gate.TCPAddr != "" {
		tcpServer = new(network.TCPServer)
		tcpServer.Addr = gate.TCPAddr
		tcpServer.MaxConnNum = gate.MaxConnNum
		tcpServer.PendingWriteNum = gate.PendingWriteNum
		tcpServer.LenMsgLen = gate.LenMsgLen
		tcpServer.MaxMsgLen = gate.MaxMsgLen
		tcpServer.LittleEndian = gate.LittleEndian
		tcpServer.NewAgent = func(conn *network.TCPConn) network.Agent {
			a := &agent{conn: conn, gate: gate}
			if gate.AgentChanRPC != nil {
				gate.AgentChanRPC.Go("NewAgent", a)
			}
			return a
		}
	}

	if wsServer != nil {
		wsServer.Start()
	}
	if tcpServer != nil {
		tcpServer.Start()
	}
	<-closeSig
	if wsServer != nil {
		wsServer.Close()
	}
	if tcpServer != nil {
		tcpServer.Close()
	}
}

func (gate *Gate) OnDestroy() {}

type agent struct {
	conn     network.Conn
	gate     *Gate
	userData interface{}
}

func (a *agent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debugf("read message: %v", err)
			break
		}

		//消息解密
		encrypt(data)

		//数据解压
		decodeData, decodeErr := compress.Decode(data)
		if decodeErr != nil {
			log.ErrorW("compress.Decode error", log.String("err", decodeErr.Error()), log.String("data", string(data)))
			break
		}
		data = decodeData

		if a.gate.Processor != nil {
			msg, err := a.gate.Processor.Unmarshal(data)
			if err != nil {
				log.Debugf("unmarshal message error: %v", err)
				break
			}
			err = a.gate.Processor.Route(msg, a)
			if err != nil {
				log.Debugf("route message error: %v", err)
				break
			}
		}
	}
}

func (a *agent) OnClose() {
	if a.gate.AgentChanRPC != nil {
		err := a.gate.AgentChanRPC.Call0("CloseAgent", a)
		if err != nil {
			log.Errorf("chanrpc error: %v", err)
		}
	}
}

func (a *agent) WriteMsg(msg interface{}) {
	if a.gate.Processor != nil {
		data, err := a.gate.Processor.Marshal(msg)
		if err != nil {
			log.Errorf("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}

		//消息加密
		for i := 0; i < len(data); i++ {

			//压缩前
			oldLen := len(data[i])

			//数据压缩
			encodeData, encodeErr := compress.Encode(data[i])
			if encodeErr != nil {
				log.ErrorW("compress.Encode error", log.String("err", encodeErr.Error()), log.String("data", string(data[i])))
				continue
			}
			data[i] = encodeData

			//压缩后
			newLen := len(data[i])

			//sizeLimit := 1024 * 10
			sizeLimit := conf.MsgSizeLimit
			if oldLen > sizeLimit {
				text := compress.Name() + "压缩比"
				name := fmt.Sprint(reflect.TypeOf(msg))
				log.WarnW(text,
					log.String("消息名", name),
					log.Int("压缩前", oldLen),
					log.Int("压缩后", newLen),
					log.Float64("压缩比", float64(newLen)/float64(oldLen)))
			}

			encrypt(data[i])
		}

		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Errorf("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}

func (a *agent) WriteRaw(msgId string, data []byte) {
	if a.gate.Processor != nil {
		//压缩前
		oldLen := len(data)

		//数据压缩
		encodeData, encodeErr := compress.Encode(data)
		if encodeErr != nil {
			log.ErrorW("compress.Encode error", log.String("err", encodeErr.Error()), log.String("data", string(data)))
			return
		}
		data = encodeData

		//压缩后
		newLen := len(data)

		//sizeLimit := 1024 * 10
		sizeLimit := conf.MsgSizeLimit
		if oldLen > sizeLimit {
			text := compress.Name() + "压缩比"
			name := fmt.Sprint(reflect.TypeOf(data))
			log.WarnW(text,
				log.String("消息名", name),
				log.Int("压缩前", oldLen),
				log.Int("压缩后", newLen),
				log.Float64("压缩比", float64(newLen)/float64(oldLen)))
		}

		encrypt(data)
	}

	err := a.conn.WriteMsg(data)
	if err != nil {
		log.Errorf("write message %v error: %v", msgId, err)
	}
}

func (a *agent) Close() {
	a.conn.Close()
}

func (a *agent) Destroy() {
	a.conn.Destroy()
}

func (a *agent) UserData() interface{} {
	return a.userData
}

func (a *agent) SetUserData(data interface{}) {
	a.userData = data
}

func (a *agent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *agent) TrueClientIP() string {
	return a.conn.TrueClientIP()
}

func (a *agent) GetHeader(key string) string {
	return a.conn.GetHeader(key)
}
