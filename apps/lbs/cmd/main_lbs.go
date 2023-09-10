package main

import (
	"im/apps/lbs/dig"
	"im/apps/lbs/internal/config"
	"im/pkg/commands"
	"im/pkg/common/xmongo"
	"im/pkg/common/xmysql"
	"im/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmongo.NewMongoClient(conf.Mongo)
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
