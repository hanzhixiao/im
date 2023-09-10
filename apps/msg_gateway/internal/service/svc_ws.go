package service

import (
	"github.com/go-playground/validator/v10"
	msg_client "im/apps/message/client"
	"im/apps/msg_gateway/internal/config"
	"im/apps/msg_gateway/internal/server/websocket/ws"
)

type WsService interface {
	Run()
	MessageCallback(msg *ws.Message)
}

type wsService struct {
	conf      *config.Config
	validate  *validator.Validate
	msgClient msg_client.MsgClient
}

func NewWsService(conf *config.Config) WsService {
	msgClient := msg_client.NewMsgClient(conf.Etcd, conf.MessageServer, conf.GrpcServer.Jaeger, conf.Name)
	return &wsService{conf: conf, validate: validator.New(), msgClient: msgClient}
}

func (s *wsService) Run() {

}
