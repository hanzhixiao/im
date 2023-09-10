package dig

import (
	"go.uber.org/dig"
	"im/apps/auth/internal/config"
	"im/apps/auth/internal/server"
	"im/apps/auth/internal/server/auth"
	"im/apps/auth/internal/service"
	"im/domain/cache"
	"im/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(auth.NewAuthServer)
	Provide(service.NewAuthService)
	Provide(repo.NewAuthRepository)
	Provide(repo.NewAvatarRepository)
	Provide(repo.NewUserRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(cache.NewAuthCache)
	Provide(cache.NewUserCache)
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
