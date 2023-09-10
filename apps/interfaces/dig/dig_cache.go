package dig

import (
	"im/domain/cache"
)

func provideCache() {
	Provide(cache.NewServerMgrCache)
}
