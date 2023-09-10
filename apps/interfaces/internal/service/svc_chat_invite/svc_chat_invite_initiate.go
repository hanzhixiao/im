package svc_chat_invite

import (
	"github.com/jinzhu/copier"
	"im/apps/interfaces/internal/dto/dto_chat_invite"
	"im/pkg/common/xlog"
	"im/pkg/proto/pb_invite"
	"im/pkg/xhttp"
)

func (s *chatInviteService) InitiateChatInvite(params *dto_chat_invite.InitiateChatInviteReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_invite.InitiateChatInviteReq)
		reply   *pb_invite.InitiateChatInviteResp
	)
	copier.Copy(reqArgs, params)
	reqArgs.InitiatorUid = uid
	reply = s.chatInviteClient.InitiateChatInvite(reqArgs)
	if reply == nil {
		resp.SetResult(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		xlog.Warn(xhttp.ERROR_CODE_HTTP_SERVICE_FAILURE, xhttp.ERROR_HTTP_SERVICE_FAILURE)
		return
	}
	if reply.Code > 0 {
		resp.SetResult(reply.Code, reply.Msg)
		xlog.Warn(reply.Code, reply.Msg)
		return
	}
	return
}
