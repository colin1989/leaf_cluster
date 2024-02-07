package main

import (
	"flag"
	"server/game"

	"github.com/name5566/leaf"
)

var ServerID int
var wsAddr string

func init() {
	flag.IntVar(&ServerID, "s", 1, "服务器ID")
	flag.StringVar(&wsAddr, "ws", "127.0.0.1:14561", "websocket 地址")
}

func main() {
	flag.Parse()

	game.SetServerID(int32(ServerID))
	game.SetWSAddr(wsAddr)
	leaf.Run(
		game.Module,
	)
}
