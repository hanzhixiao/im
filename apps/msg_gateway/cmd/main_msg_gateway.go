package main

import (
	"im/apps/msg_gateway/dig"
	"im/apps/msg_gateway/internal/config"
	"im/pkg/commands"
	"im/pkg/common/xredis"
	"runtime"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
