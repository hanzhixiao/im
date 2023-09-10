package dig

import (
	"im/apps/interfaces/internal/service/svc_red_env_receive"
)

func provideRedEnvReceive() {
	Provide(svc_red_env_receive.NewRedEnvReceiveService)
}
