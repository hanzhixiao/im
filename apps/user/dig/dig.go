package dig

import (
	"go.uber.org/dig"
	"im/apps/user/internal/config"
	"im/apps/user/internal/server"
	"im/apps/user/internal/server/user"
	"im/apps/user/internal/service"
	"im/domain/cache"
	"im/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(user.NewUserServer)
	Provide(service.NewUserService)
	Provide(repo.NewUserRepository)
	Provide(repo.NewAvatarRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(cache.NewChatMemberCache)
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
