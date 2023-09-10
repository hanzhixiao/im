package dig

import (
	"go.uber.org/dig"
	"im/apps/msg_hot/internal/config"
	"im/apps/msg_hot/internal/server"
	"im/apps/msg_hot/internal/server/msg_hot"
	"im/apps/msg_hot/internal/service"
	"im/domain/mrepo"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	Provide(server.NewServer)
	Provide(mrepo.NewMessageHotRepository)
	Provide(msg_hot.NewMessageHotServer)
	Provide(service.NewMessageHotService)
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
