package main

import (
	"client/internal"
	"flag"
	"fmt"

	"github.com/name5566/leaf"
)

func init() {
	flag.IntVar(&internal.ServerID, "s", 1, "连接游戏服ID")
	flag.StringVar(&internal.GateAddr, "addr", "127.0.0.1:13561", "连接游戏服ID")
}

func main() {
	flag.Parse()

	fmt.Println("client run...")
	leaf.Run(Module)
}
