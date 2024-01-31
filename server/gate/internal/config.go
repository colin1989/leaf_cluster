package internal

var (
	serverID = int32(1)
	wsAddr   = "127.0.0.1:13561"
)

func SetServerID(id int32) {
	serverID = id
}

func SetWSAddr(addr string) {
	wsAddr = addr
}
