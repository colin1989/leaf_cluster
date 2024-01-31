package main

import (
	"flag"
	"server/gate"

	"github.com/name5566/leaf"
)

var ServerID int
var wsAddr string

func init() {
	flag.IntVar(&ServerID, "s", 1, "服务器ID")
	flag.StringVar(&wsAddr, "ws", "127.0.0.1:13563", "websocket 地址")
}
func main() {
	flag.Parse()

	gate.SetServerID(int32(ServerID))
	gate.SetWSAddr(wsAddr)
	leaf.Run(
		gate.Module,
		gate.GateModule,
	)
}
