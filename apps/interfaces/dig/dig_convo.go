package dig

import (
	"im/apps/interfaces/internal/service/svc_convo"
)

func provideConvo() {
	Provide(svc_convo.NewConvoService)
}
