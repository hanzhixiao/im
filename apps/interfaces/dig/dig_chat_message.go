package dig

import (
	"im/apps/interfaces/internal/service/svc_chat_msg"
)

func provideChatMessage() {
	Provide(svc_chat_msg.NewChatMessageService)
}
