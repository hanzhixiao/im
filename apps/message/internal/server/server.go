package server

import (
	"im/apps/message/internal/server/message"
	"im/pkg/commands"
)

type server struct {
	messageServer message.MessageServer
}

func NewServer(messageServer message.MessageServer) commands.MainInstance {
	return &server{messageServer: messageServer}
}

func (s *server) Initialize() (err error) {
	return
}

func (s *server) RunLoop() {
	s.messageServer.Run()
}

func (s *server) Destroy() {

}
