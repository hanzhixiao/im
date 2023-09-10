package dig

import (
	"im/apps/interfaces/internal/service/svc_chat_invite"
)

func provideChatInvite() {
	Provide(svc_chat_invite.NewChatInviteService)
}
