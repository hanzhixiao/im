package main

import (
	"im/apps/msg_hot/dig"
	"im/apps/msg_hot/internal/config"
	"im/pkg/commands"
	"im/pkg/common/xmongo"
	"im/pkg/common/xredis"
)

// 弃用
func init() {
	conf := config.GetConfig()
	xmongo.NewMongoClient(conf.Mongo)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
