package gateway

import (
	"im/apps/msg_gateway/internal/logic"
	"im/pkg/proto/pb_enum"
	"im/pkg/proto/pb_gw"
)

func (s *gatewayServer) sendInviteMessage(req *pb_gw.SendTopicMessageReq, resp *pb_gw.SendTopicMessageResp) {
	var (
		message *pb_gw.SendMessage
		buf     []byte
		err     error
	)
	message, buf, err = AssemblyMessage(req, req.SenderId, req.SenderPlatform, pb_enum.MESSAGE_TYPE_NEW, resp)
	if err != nil {
		return
	}
	logic.SendMessage(message, buf, s.wsServer.SendMessage, s.SendCloudMessage)
	return
}
