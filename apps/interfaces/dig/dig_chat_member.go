package dig

import (
	"im/apps/interfaces/internal/service/svc_chat_member"
)

func provideChatMember() {
	Provide(svc_chat_member.NewChatMemberService)
}
