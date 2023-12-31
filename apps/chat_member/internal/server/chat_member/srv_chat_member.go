package chat_member

import (
	"google.golang.org/grpc"
	"im/apps/chat_member/internal/config"
	"im/apps/chat_member/internal/service"
	"im/pkg/common/xgrpc"
	"im/pkg/proto/pb_chat_member"
	"io"
)

type ChatMemberServer interface {
	Run()
}

type chatMemberServer struct {
	pb_chat_member.UnimplementedChatMemberServer
	cfg               *config.Config
	grpcServer        *xgrpc.GrpcServer
	chatMemberService service.ChatMemberService
}

func NewChatMemberServer(cfg *config.Config, chatMemberService service.ChatMemberService) ChatMemberServer {
	srv := &chatMemberServer{cfg: cfg, chatMemberService: chatMemberService}
	return srv
}

func (s *chatMemberServer) Run() {
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

	pb_chat_member.RegisterChatMemberServer(srv, s)
	s.grpcServer = xgrpc.NewGrpcServer(s.cfg.GrpcServer, s.cfg.Etcd)
	s.grpcServer.RunServer(srv)
}
