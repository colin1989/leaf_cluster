package config

import "time"

type WSConfig struct {
	//masterAddr      string
	WSAddr          string
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	HTTPTimeout     time.Duration
}

func DefaultWSConfig(addr string) WSConfig {
	return WSConfig{
		WSAddr:          addr,
		MaxConnNum:      20000,
		PendingWriteNum: 20000,
		MaxMsgLen:       40960,
		HTTPTimeout:     10 * time.Second,
	}
}
