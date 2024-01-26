package main

import (
	"server/game"

	"github.com/name5566/leaf"
)

func main() {
	leaf.Run(
		game.Module,
		game.GateModule,
	)
}
