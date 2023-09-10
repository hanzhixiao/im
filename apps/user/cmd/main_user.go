package main

import (
	"im/apps/user/dig"
	"im/apps/user/internal/config"
	"im/pkg/commands"
	"im/pkg/common/xes"
	"im/pkg/common/xmysql"
	"im/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
	xes.NewElasticsearchClient(conf.Elasticsearch)
}

func main() {
	dig.Invoke(func(srv commands.MainInstance) {
		commands.Run(srv)
	})
}
