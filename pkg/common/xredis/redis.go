package xredis

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"im/pkg/common/xlog"
	"im/pkg/conf"
)

var (
	Cli *RedisClient
)

type RedisClient struct {
	Client   redis.UniversalClient
	RedsSync *redsync.Redsync
	Prefix   string
}

func NewRedisClient(cfg *conf.Redis) redis.UniversalClient {
	var (
		client   redis.UniversalClient
		pool     redsyncredis.Pool
		redsSync *redsync.Redsync
		err      error
	)
	// 集群redis
	if len(cfg.Address) > 1 {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: cfg.Address,
		})
	} else {
		client = redis.NewClient(&redis.Options{Addr: cfg.Address[0]})
	}
	// 判断是否能够链接到redis
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		xlog.Error(err.Error())
	}
	// redis 锁
	pool = goredis.NewPool(client)
	redsSync = redsync.New(pool)

	Cli = &RedisClient{client, redsSync, cfg.Prefix}
	return client
}
