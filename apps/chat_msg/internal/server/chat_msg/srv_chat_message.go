package chat_msg

import (
	"google.golang.org/grpc"
	"im/apps/chat_msg/internal/config"
	"im/apps/chat_msg/internal/service"
	"im/pkg/common/xgrpc"
	"im/pkg/proto/pb_chat_msg"
	"io"
)

type ChatMessageServer interface {
	Run()
}

type chatMessageServer struct {
	pb_chat_msg.UnimplementedChatMessageServer
	cfg                *config.Config
	grpcServer         *xgrpc.GrpcServer
	chatMessageService service.ChatMessageService
}

func NewChatMessageServer(cfg *config.Config, chatMessageService service.ChatMessageService) ChatMessageServer {
	return &chatMessageServer{cfg: cfg, chatMessageService: chatMessageService}
}

func (s *chatMessageServer) Run() {
	var (
		srv    *grpc.Server
		closer io.Closer
	)
	srv, closer = xgrpc.NewServer(s.cfg.GrpcServer)
	defer func() {
		if closer != nil {
			closer.Close()
		}
	}()

	pb_chat_msg.RegisterChatMessageServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
