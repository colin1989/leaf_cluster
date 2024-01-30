package main

import (
	"client/msg"
	"message"
)

func init() {
	// 消息路由到Game server
	msg.JSONProcessor.SetRouter(&message.Greeting{}, ChanRPC)
	msg.JSONProcessor.SetRouter(&message.Login{}, ChanRPC)
}
