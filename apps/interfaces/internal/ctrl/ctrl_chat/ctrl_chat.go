package ctrl_chat

import (
	"im/apps/interfaces/internal/service/svc_chat"
)

type ChatCtrl struct {
	chatService svc_chat.ChatService
}

func NewChatCtrl(chatService svc_chat.ChatService) *ChatCtrl {
	return &ChatCtrl{chatService: chatService}
}
