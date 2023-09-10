package main

import (
	"im/apps/chat_member/dig"
	"im/apps/chat_member/internal/config"
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
