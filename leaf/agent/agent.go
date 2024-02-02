package agent

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	WriteRaw(data []byte)
	Close()
	Destroy()
	ID() int64
	UserData() interface{}
	SetUserData(data interface{})
	RemoteAddr() net.Addr
	TrueClientIP() string
	GetHeader(string) string
}
