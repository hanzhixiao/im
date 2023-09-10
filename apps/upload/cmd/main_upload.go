package main

import (
	"im/apps/upload/internal/config"
	"im/apps/upload/internal/server"
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
