package gate

import (
	"fmt"
	"net"
	"reflect"
	"sync/atomic"

	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/cluster/protos"
	"github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/util/compress"
)

type agent struct {
	conn        network.Conn
	gate        *Gate
	clusterGate cluster.ClusterNode
	id          int64
	uid         int64
	userData    interface{}
}

var agentID = int64(0)

func newAgent(conn network.Conn, gate *Gate) (network.Agent, int64) {
	aid := atomic.LoadInt64(&agentID)
	atomic.AddInt64(&agentID, 1)
	return &agent{
		conn:        conn,
		gate:        gate,
		clusterGate: cluster.GetNode(),
		id:          aid,
		userData:    nil,
	}, agentID
}

const HeadLength = 8

func (a *agent) Run() {
	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debugf("read message: %v", err)
			break
		}

		// 读取消息头
		if len(data) < HeadLength {
			log.Error("head length less then 8")
			break
		}

		head := data[:8]
		_ = head
		data = data[8:]

		//消息解密
		encrypt(data)

		//数据解压
		decodeData, decodeErr := compress.Decode(data)
		if decodeErr != nil {
			log.ErrorW("compress.Decode error", log.String("err", decodeErr.Error()), log.String("data", string(data)))
			break
		}
		data = decodeData

		if a.forwardGate(data) {
			continue
		}

		if a.gate.Processor != nil {
			msg, err := a.gate.Processor.Unmarshal(data)
			if err != nil {
				log.Debugf("unmarshal message error: %v", err)
				break
			}
			err = a.gate.Processor.Route(msg, a, a.userData)
			if err != nil {
				log.Debugf("route message error: %v", err)
				break
			}
		}
	}
}

func (a *agent) forwardGate(data []byte) bool {
	sd, ok := a.UserData().(*protos.SessionData)
	if !ok {
		return false
	}
	if sd.UId == 0 {
		return false
	}

	m := &protos.Msg{
		Id:    0,
		Route: "",
		Data:  [][]byte{data},
		Type:  protos.MsgType_MsgLogin,
	}

	req := &protos.Request{
		Session: sd,
		Msg:     m,
		Server:  a.clusterGate.Server(),
	}

	rpcServer, ok := a.clusterGate.(cluster.ClusterRPC)
	if !ok {
		log.Error("node does not implement ClusterRPC")
		return false
	}

	if err := rpcServer.RPCTo(sd.SId, req); err != nil {
		log.Error("RPCTo error", log.Int32("sid", sd.SId),
			log.FieldErr(err))
		return false
	}
	return true
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
			data[i] = append([]byte{0, 0, 0, 0, 0, 0, 0, 0}, data[i]...)
		}

		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Errorf("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}

func (a *agent) WriteRaw(data []byte) {
	//消息加密

	//压缩前
	oldLen := len(data)

	//数据压缩
	encodeData, encodeErr := compress.Encode(data)
	if encodeErr != nil {
		log.ErrorW("compress.Encode error",
			log.String("err", encodeErr.Error()), log.String("data", string(data)))
		return
	}
	data = encodeData

	//压缩后
	newLen := len(data)

	//sizeLimit := 1024 * 10
	sizeLimit := conf.MsgSizeLimit
	if oldLen > sizeLimit {
		text := compress.Name() + "压缩比"
		name := "WriteRaw" //fmt.Sprint(reflect.TypeOf(msg))
		log.WarnW(text,
			log.String("消息名", name),
			log.Int("压缩前", oldLen),
			log.Int("压缩后", newLen),
			log.Float64("压缩比", float64(newLen)/float64(oldLen)))
	}

	encrypt(data)
	data = append([]byte{0, 0, 0, 0, 0, 0, 0, 0}, data...)

	err := a.conn.WriteMsg(data)
	if err != nil {
		log.Errorf("write message %v error: %v", "WriteRaw", err)
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

func (a *agent) ID() int64 {
	return a.id
}
