package main

import (
	"client/internal"
	"client/msg"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func init() {
	// 消息路由到Game server
	msg.JSONProcessor.SetRouter(&msg.Greeting{}, ChanRPC)
}
