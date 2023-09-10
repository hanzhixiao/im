package dig

import (
	"go.uber.org/dig"
	"im/apps/lbs/internal/config"
	"im/apps/lbs/internal/server"
	"im/apps/lbs/internal/server/lbs"
	"im/apps/lbs/internal/service"
	"im/domain/cache"
	"im/domain/mrepo"
	"im/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(lbs.NewLbsServer)
	Provide(service.NewLbsService)
	Provide(mrepo.NewLbsRepository)
	Provide(repo.NewUserRepository)
	Provide(repo.NewUserLocationRepository)
	Provide(cache.NewUserCache)
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
