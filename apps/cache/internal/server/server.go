package server

import (
	srv_cache "im/apps/cache/internal/server/cache"
	"im/pkg/commands"
)

type server struct {
	cacheServer srv_cache.CacheServer
}

func NewServer(cacheServer srv_cache.CacheServer) commands.MainInstance {
	return &server{cacheServer: cacheServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {

}

func (s *server) Destroy() {

}
