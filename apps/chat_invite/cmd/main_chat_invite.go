package main

import (
	"im/apps/chat_invite/dig"
	"im/apps/chat_invite/internal/config"
	"im/pkg/commands"
	"im/pkg/common/xmysql"
	"im/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
