package server

import (
	"im/apps/msg_gateway/internal/server/gateway"
	"im/pkg/commands"
)

type server struct {
	gatewayServer gateway.GatewayServer
}

func NewServer(gatewayServer gateway.GatewayServer) commands.MainInstance {
	return &server{gatewayServer: gatewayServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.gatewayServer.Run()
}

func (s *server) Destroy() {

}
