package main

import (
	"im/apps/interfaces/internal/config"
	_ "im/apps/interfaces/internal/config"
	"im/apps/interfaces/internal/server"
	"im/pkg/commands"
	"im/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	commands.Run(server.NewServer())
}
