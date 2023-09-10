package dig

import (
	"go.uber.org/dig"
	"im/apps/chat_member/internal/config"
	"im/apps/chat_member/internal/server"
	"im/apps/chat_member/internal/server/chat_member"
	"im/apps/chat_member/internal/service"
	"im/domain/cache"
	"im/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(chat_member.NewChatMemberServer)
	Provide(service.NewChatMemberService)
	Provide(repo.NewChatMemberRepository)
	Provide(repo.NewUserRepository)
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
