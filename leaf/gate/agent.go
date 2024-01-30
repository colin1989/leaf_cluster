package gate

import "net"

type Agent interface {
	WriteMsg(msg interface{})
	WriteRaw(msgId string, data []byte)
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
	RemoteAddr() net.Addr
	TrueClientIP() string
	GetHeader(string) string
}
