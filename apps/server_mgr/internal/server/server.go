package server

import (
	"im/apps/server_mgr/internal/server/server_mgr"
	"im/pkg/commands"
)

type server struct {
	serverMgrServer server_mgr.ServerMgrServer
}

func NewServer(serverMgrServer server_mgr.ServerMgrServer) commands.MainInstance {
	return &server{serverMgrServer: serverMgrServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.serverMgrServer.Run()
}

func (s *server) Destroy() {

}
