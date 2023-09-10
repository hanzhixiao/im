package main

import (
	"im/apps/server_mgr/dig"
	"im/apps/server_mgr/internal/config"
	"im/pkg/commands"
	"im/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
