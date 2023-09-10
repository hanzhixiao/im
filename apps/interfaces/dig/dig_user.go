package dig

import (
	"im/apps/interfaces/internal/service/svc_user"
)

func provideUser() {
	Provide(svc_user.NewUserService)
}
