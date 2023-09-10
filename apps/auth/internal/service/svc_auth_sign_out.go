package service

import (
	"context"
	"im/pkg/common/xlog"
	"im/pkg/proto/pb_auth"
	"im/pkg/proto/pb_chat_member"
)

func (s *authService) SignOut(ctx context.Context, req *pb_auth.SignOutReq) (resp *pb_auth.SignOutResp, _ error) {
	resp = new(pb_auth.SignOutResp)
	var (
		err       error
		onOffResp *pb_chat_member.ChatMemberOnOffLineResp
	)
	onOffResp = s.chatMemberOnOffLine(req.Uid, 0, req.Platform)
	if onOffResp == nil {
		resp.Set(ERROR_CODE_AUTH_GRPC_SERVICE_FAILURE, ERROR_AUTH_GRPC_SERVICE_FAILURE)
		xlog.Warn(ERROR_CODE_AUTH_GRPC_SERVICE_FAILURE, ERROR_AUTH_GRPC_SERVICE_FAILURE)
		return
	}
	if onOffResp.Code > 0 {
		resp.Set(onOffResp.Code, onOffResp.Msg)
		xlog.Warn(onOffResp.Code, onOffResp.Msg)
		return
	}
	err = s.userCache.SignOut(req.Uid, req.Platform)
	if err != nil {
		resp.Set(ERROR_CODE_AUTH_LOGOUT_FAILED, ERROR_AUTH_LOGOUT_FAILED)
		xlog.Warn(ERROR_CODE_AUTH_LOGOUT_FAILED, ERROR_AUTH_LOGOUT_FAILED, err.Error())
		return
	}
	return
}
