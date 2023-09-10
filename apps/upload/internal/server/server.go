package server

import (
	"flag"
	"im/apps/upload/internal/config"
	"im/apps/upload/internal/router"
	"im/pkg/commands"
	"im/pkg/common/xetcd"
	"im/pkg/common/xgin"
	"im/pkg/common/xlog"
	"im/pkg/utils"
)

type server struct {
	ginServer *xgin.GinServer
	cfg       *config.Config
}

func NewServer() commands.MainInstance {
	return &server{cfg: config.NewConfig()}
}

func (s *server) Initialize() (err error) {
	flag.Parse()
	s.ginServer = xgin.NewGinServer()
	router.Register(s.ginServer.Engine)
	return
}

func (s *server) RunLoop() {
	err := xetcd.RegisterEtcd(s.cfg.Etcd.Schema, s.cfg.Etcd.Endpoints, utils.GetServerIP(), s.cfg.Port, s.cfg.Name, xetcd.TIME_TO_LIVE)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	s.ginServer.Run(s.cfg.Port)
}

func (s *server) Destroy() {

}
