package dig

import (
	"im/apps/interfaces/internal/service/svc_lbs"
)

func provideLbs() {
	Provide(svc_lbs.NewLbsService)
}
