package dig

import (
	"go.uber.org/dig"
	"im/apps/server_mgr/internal/config"
	"im/apps/server_mgr/internal/server"
	"im/apps/server_mgr/internal/server/server_mgr"
	"im/apps/server_mgr/internal/service"
	"im/domain/cache"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(server_mgr.NewServerMgrServer)
	Provide(service.NewServerMgrService)
	Provide(cache.NewServerMgrCache)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}

func Provide(constructor interface{}, opts ...dig.ProvideOption) {
	err := container.Provide(constructor)
	if err != nil {
		log.Panic(err)
	}
}
