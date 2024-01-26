package gate

import "net"

type Agent interface {
	WriteMsg(msg interface{})
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
	RemoteAddr() net.Addr
	TrueClientIP() string
	GetHeader(string) string
}
