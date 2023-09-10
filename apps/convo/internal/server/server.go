package server

import (
	"im/apps/convo/internal/server/convo"
	"im/pkg/commands"
)

type server struct {
	convoServer convo.ConvoServer
}

func NewServer(convoServer convo.ConvoServer) commands.MainInstance {
	return &server{convoServer: convoServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.convoServer.Run()
}

func (s *server) Destroy() {

}
