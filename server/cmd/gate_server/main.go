package main

import (
	"server/gate"

	"github.com/name5566/leaf"
)

func main() {
	leaf.Run(
		gate.Module,
		gate.GateModule,
	)
}
