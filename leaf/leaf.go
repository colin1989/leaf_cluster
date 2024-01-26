package leaf

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/name5566/leaf/cluster"
	"github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/console"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/log/zap"
	"github.com/name5566/leaf/module"
)

func Run(mods ...module.Module) {
	// logger
	config := &log.Config{
		Debug: conf.LogDebug,
		Level: conf.LogLevel,
		Net:   "udp",
		Addr:  conf.LogAddr,
		Tag:   conf.LogTag,
	}
	log.SetLogger(zap.NewLogger(config.BuildOption()...))

	log.Infof("dhf %v starting up", version)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// cluster
	cluster.Init()

	// console
	console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	sig := <-c
	log.Infof("dhf closing down (signal: %v)", sig)
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}
