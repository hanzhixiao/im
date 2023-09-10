package convo

import (
	"google.golang.org/grpc"
	"im/apps/convo/internal/config"
	"im/apps/convo/internal/service"
	"im/pkg/common/xgrpc"
	"im/pkg/common/xmonitor"
	"im/pkg/proto/pb_convo"
	"io"
)

type ConvoServer interface {
	Run()
}

// conversation
type convoServer struct {
	pb_convo.UnimplementedConvoServer
	cfg          *config.Config
	convoService service.ConvoService
	grpcServer   *xgrpc.GrpcServer
}

func NewConvoServer(cfg *config.Config, convoService service.ConvoService) ConvoServer {
	srv := &convoServer{cfg: cfg, convoService: convoService}
	return srv
}

func (s *convoServer) Run() {
	var (
		srv    *grpc.Server
		closer io.Closer
	)
	xmonitor.RunMonitor(s.cfg.Monitor.Port)

	srv, closer = xgrpc.NewServer(s.cfg.GrpcServer)
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()

	pb_convo.RegisterConvoServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
