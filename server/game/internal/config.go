package internal

var (
	Key      = "FK44r1uPz9x7z8Ph75LIEI7nhp6H40R8"
	serverID = int32(1)
	wsAddr   = "127.0.0.1:14561"
)

func SetServerID(id int32) {
	serverID = id
}

func SetWSAddr(addr string) {
	wsAddr = addr
}
