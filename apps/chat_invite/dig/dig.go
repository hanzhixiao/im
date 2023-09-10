package dig

import (
	"go.uber.org/dig"
	"im/apps/chat_invite/internal/config"
	"im/apps/chat_invite/internal/server"
	"im/apps/chat_invite/internal/server/chat_invite"
	"im/apps/chat_invite/internal/service"
	"im/domain/cache"
	"im/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(chat_invite.NewChatInviteServer)
	Provide(service.NewChatInviteService)
	Provide(repo.NewChatInviteRepository)
	Provide(repo.NewUserRepository)
	Provide(repo.NewAvatarRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(repo.NewChatRepository)
	Provide(cache.NewChatCache)
	Provide(cache.NewChatMessageCache)
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
