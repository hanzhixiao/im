package dig

import (
	"im/apps/interfaces/internal/service/svc_red_env"
)

func provideRedEnv() {
	Provide(svc_red_env.NewRedEnvService)
}
