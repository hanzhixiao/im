package server_mgr

import (
	"im/apps/server_mgr/internal/config"
	"im/apps/server_mgr/internal/service"
)

type ServerMgrServer interface {
	Run()
}

type serverMgrServer struct {
	cfg              *config.Config
	serverMgrService service.ServerMgrService
}

func NewServerMgrServer(cfg *config.Config, serverMgrService service.ServerMgrService) ServerMgrServer {
	return &serverMgrServer{cfg: cfg, serverMgrService: serverMgrService}
}

func (s *serverMgrServer) Run() {
	s.serverMgrService.Run()
}
