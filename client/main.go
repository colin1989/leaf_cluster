package main

import (
	"client/internal"
	"flag"
	"fmt"

	"github.com/name5566/leaf"
)

var ServerID int

func init() {
	flag.IntVar(&ServerID, "s", 1, "用户名,默认为空")
}

func main() {
	flag.Parse()

	internal.SetServerID(int32(ServerID))
	fmt.Println("client run...")
	leaf.Run(Module)
}
