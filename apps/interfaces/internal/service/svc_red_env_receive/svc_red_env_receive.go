package svc_red_env_receive

import (
	"im/apps/interfaces/internal/config"
	"im/apps/interfaces/internal/dto/dto_red_env_receive"
	red_env_receive_client "im/apps/red_env_receive/client"
	"im/pkg/xhttp"
)

type RedEnvReceiveService interface {
	GrabRedEnvelope(params *dto_red_env_receive.GrabRedEnvelopeReq, uid int64) (resp *xhttp.Resp)
	OpenRedEnvelope(params *dto_red_env_receive.OpenRedEnvelopeReq, uid int64) (resp *xhttp.Resp)
}

type redEnvReceiveService struct {
	redEnvReceiveClient red_env_receive_client.RedEnvReceiveClient
}

func NewRedEnvReceiveService(conf *config.Config) RedEnvReceiveService {
	redEnvReceiveClient := red_env_receive_client.NewRedEnvReceiveClient(conf.Etcd, conf.RedEnvReceiveServer, conf.Jaeger, conf.Name)
	return &redEnvReceiveService{redEnvReceiveClient: redEnvReceiveClient}
}
