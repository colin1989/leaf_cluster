package main

import (
	"server/world"

	"github.com/name5566/leaf"
)

func main() {
	leaf.Run(
		world.Module,
		world.GateModule,
	)
}
