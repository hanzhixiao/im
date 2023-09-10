package dig

import (
	"go.uber.org/dig"
	"im/apps/message/internal/config"
	"im/apps/message/internal/server"
	"im/apps/message/internal/server/message"
	"im/apps/message/internal/service"
	"im/domain/cache"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(message.NewMessageServer)
	Provide(service.NewMessageService)
	Provide(cache.NewChatMemberCache)
	Provide(cache.NewChatMessageCache)

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
