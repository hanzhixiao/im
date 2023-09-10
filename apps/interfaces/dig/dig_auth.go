package dig

import (
	"im/apps/interfaces/internal/service/svc_auth"
)

func provideAuth() {
	Provide(svc_auth.NewAuthService)
}
