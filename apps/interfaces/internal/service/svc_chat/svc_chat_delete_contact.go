package svc_chat

import (
	"github.com/jinzhu/copier"
	"im/apps/interfaces/internal/dto/dto_chat"
	"im/pkg/common/xlog"
	"im/pkg/proto/pb_chat"
	"im/pkg/xhttp"
)

func (s *chatService) DeleteContact(params *dto_chat.DeleteContactReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		reqArgs = new(pb_chat.DeleteContactReq)
		reply   *pb_chat.DeleteContactResp
	)
	copier.Copy(reqArgs, params)
	reqArgs.Uid = uid
	reply = s.chatClient.DeleteContact(reqArgs)
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
