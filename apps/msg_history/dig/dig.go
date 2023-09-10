package dig

import (
	"go.uber.org/dig"
	"im/apps/msg_history/internal/config"
	"im/apps/msg_history/internal/server"
	"im/apps/msg_history/internal/server/msg_history"
	"im/apps/msg_history/internal/service"
	"im/domain/cache"
	"im/domain/repo"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(msg_history.NewMessageHistoryServer)
	Provide(service.NewMessageHistoryService)
	Provide(repo.NewChatMessageRepository)
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
