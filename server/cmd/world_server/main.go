package main

import (
	"github.com/name5566/leaf"
	"server/world"
)

func main() {
	leaf.Run(
		world.Module,
	)
}
