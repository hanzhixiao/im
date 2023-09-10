package svc_red_env

import (
	"im/apps/interfaces/internal/config"
	"im/apps/interfaces/internal/dto/dto_red_env"
	red_env_client "im/apps/red_env/client"
	"im/pkg/xhttp"
)

type RedEnvService interface {
	GiveRedEnvelope(params *dto_red_env.GiveRedEnvelopeReq, uid int64) (resp *xhttp.Resp)
}

type redEnvService struct {
	redEnvClient red_env_client.RedEnvClient
}

func NewRedEnvService(conf *config.Config) RedEnvService {
	redEnvClient := red_env_client.NewRedEnvClient(conf.Etcd, conf.RedEnvServer, conf.Jaeger, conf.Name)
	return &redEnvService{redEnvClient: redEnvClient}
}
