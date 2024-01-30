package main

import (
	"flag"
	"server/game"

	"github.com/name5566/leaf"
)

func main() {
	var ServerID int
	var wsAddr string
	flag.IntVar(&ServerID, "s", 1, "用户名,默认为空")
	flag.StringVar(&wsAddr, "ws", "127.0.0.1:14561", "用户名,默认为空")

	game.SetServerID(ServerID)
	game.SetWSAddr(wsAddr)
	leaf.Run(
		game.Module,
		game.GateModule,
	)
}
