package dig

import (
	"go.uber.org/dig"
	"im/apps/chat_msg/internal/config"
	"im/apps/chat_msg/internal/server"
	"im/apps/chat_msg/internal/server/chat_msg"
	"im/apps/chat_msg/internal/service"
	"im/domain/cache"
	"im/domain/mrepo"
	"im/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(chat_msg.NewChatMessageServer)
	Provide(service.NewChatMessageService)
	Provide(repo.NewChatMessageRepository)
	Provide(repo.NewChatMemberRepository)
	Provide(mrepo.NewMessageHotRepository)
	Provide(cache.NewChatMessageCache)
	Provide(cache.NewChatMemberCache)
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
