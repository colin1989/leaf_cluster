package session

import (
	"fmt"
	"net"
	"reflect"

	"github.com/name5566/leaf/cluster/protos"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
)

type Session struct {
	//cluster  facade.ClusterNode
	conn      network.Conn
	chanRPC   *chanrpc.Server
	processor network.Processor
	userData  interface{}
	server    *protos.Server
}

func NewSession(conn network.Conn, p network.Processor, chanRPC *chanrpc.Server) *Session {
	return &Session{
		conn:      conn,
		chanRPC:   chanRPC,
		processor: p,
	}
}

func (s *Session) Run() {
	for {
		data, err := s.conn.ReadMsg()
		if err != nil {
			log.Debugf("read message: %v", err)
			break
		}

		if s.processor != nil {
			msg, err := s.processor.Unmarshal(data)
			if err != nil {
				log.Debugf("unmarshal message error: %v", err)
				break
			}
			err = s.processor.Route(msg, s, nil)
			if err != nil {
				log.Debugf("route message error: %v", err)
				break
			}
		}
	}
}

func (s *Session) OnClose() {
	if s.chanRPC != nil {
		err := s.chanRPC.Call0("CloseSession", s)
		if err != nil {
			log.Errorf("chanrpc error: %v", err)
		}
	}
}

func (s *Session) WriteMsg(msg interface{}) {
	data, err := s.processor.Marshal(msg)
	if err != nil {
		log.Errorf("marshal message %v error: %v", reflect.TypeOf(msg), err)
		return
	}

	err = s.conn.WriteMsg(data...)
	if err != nil {
		log.Errorf("write message %v error: %v", msg, err)
	}
}

func (s *Session) WriteRaw(data []byte) {

}

func (s *Session) WriteResponse(msg interface{}, sd *protos.SessionData) {
	data, err := s.processor.Marshal(msg)
	if err != nil {
		log.Errorf("marshal message %v error: %v", reflect.TypeOf(msg), err)
		return
	}

	//err = s.conn.WriteMsg(data...)
	//if err != nil {
	//	log.Errorf("write message %v error: %v", msg, err)
	//}

	m := &protos.Msg{
		Id:    0,
		Route: "",
		Data:  data,
		Type:  protos.MsgType_MsgLogin,
	}

	req := &protos.Response{
		Session: sd,
		Msg:     m,
	}

	reqData, err := s.processor.Marshal(req)
	if err != nil {
		log.Errorf("marshal message %v error: %v", reflect.TypeOf(msg), err)
		return
	}

	err = s.conn.WriteMsg(reqData...)
	if err != nil {
		log.Errorf("write message %v error: %v", msg, err)
	}
}

func (s *Session) Close() {
	s.conn.Close()
}

func (s *Session) Destroy() {
	s.conn.Destroy()
}

func (s *Session) Kick(uid, agentId int64) error {
	if uid == 0 || agentId == 0 {
		return fmt.Errorf("玩家 Id / agentId 错误")
	}
	kick := &protos.Kick{
		AgentId: agentId,
		UId:     uid,
	}
	data, err := s.processor.Marshal(kick)
	if err != nil {
		return fmt.Errorf("marshal message kick error: %v", err)
	}

	err = s.conn.WriteMsg(data...)
	if err != nil {
		log.Errorf("write message %v error: %v", data, err)
	}
	return nil
}

func (s *Session) Bind(aid, userID int64, sid int32) error {
	if s.server == nil {
		return fmt.Errorf("session 未绑定服务器数据")
	}
	if !s.server.CheckType(protos.ServerType_Gate) {
		return fmt.Errorf("session 未绑定网关数据")
	}
	s.WriteMsg(&protos.Bind{
		AgentId: aid,
		UId:     userID,
		SId:     sid,
	})
	return nil
}

func (s *Session) ID() int64 {
	log.Errorf("Session 并不会保存数据")
	return 0
}

func (s *Session) Server() *protos.Server {
	return s.server
}

func (s *Session) SetServer(server *protos.Server) {
	s.server = server
}

func (s *Session) UserData() interface{} {
	return s.userData
}

func (s *Session) SetUserData(data interface{}) {
	s.userData = data
}

func (s *Session) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}

func (s *Session) TrueClientIP() string {
	return s.conn.TrueClientIP()
}

func (s *Session) GetHeader(key string) string {
	return s.conn.GetHeader(key)
}
