package dig

import (
	"go.uber.org/dig"
	"im/apps/upload/internal/config"
	"log"
)

var container = dig.New()

func init() {
	Provide(config.NewConfig)
	provideUpload()
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
