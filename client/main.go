package main

import (
	"flag"
	"fmt"

	"github.com/name5566/leaf"
)

var ServerID int

func init() {
	flag.IntVar(&ServerID, "s", 1, "连接游戏服ID")
}

func main() {
	flag.Parse()

	fmt.Println("client run...")
	leaf.Run(Module)
}
