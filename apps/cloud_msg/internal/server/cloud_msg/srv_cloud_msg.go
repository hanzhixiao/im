package cloud_msg

import (
	"im/apps/cloud_msg/internal/config"
	"im/apps/cloud_msg/internal/service"
)

type CloudMessageServer interface {
	Run()
}

type cloudMessageServer struct {
	conf                *config.Config
	cloudMessageService service.CloudMessageService
}

func NewCloudMessageServer(conf *config.Config, cloudMessageService service.CloudMessageService) CloudMessageServer {
	return &cloudMessageServer{conf: conf, cloudMessageService: cloudMessageService}
}

func (s *cloudMessageServer) Run() {

}
