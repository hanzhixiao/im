package svc_chat_member

import (
	chat_member_client "im/apps/chat_member/client"
	"im/apps/interfaces/internal/config"
	"im/apps/interfaces/internal/dto/dto_chat_member"
	"im/pkg/xhttp"
)

type ChatMemberService interface {
	ChatMemberList(params *dto_chat_member.ChatMemberListReq) (resp *xhttp.Resp)
	ContactList(params *dto_chat_member.ContactListReq, uid int64) (resp *xhttp.Resp)
	GroupChatList(params *dto_chat_member.GroupChatListReq, uid int64) (resp *xhttp.Resp)
}

type chatMemberService struct {
	chatMemberClient chat_member_client.ChatMemberClient
}

func NewChatMemberService(conf *config.Config) ChatMemberService {
	chatMemberClient := chat_member_client.NewChatMemberClient(conf.Etcd, conf.ChatMemberServer, conf.Jaeger, conf.Name)
	return &chatMemberService{chatMemberClient: chatMemberClient}
}
