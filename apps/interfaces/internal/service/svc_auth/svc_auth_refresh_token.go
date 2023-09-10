package svc_auth

import (
	"im/apps/interfaces/internal/dto/dto_auth"
	"im/pkg/common/xlog"
	"im/pkg/proto/pb_auth"
	"im/pkg/xhttp"
)

func (s *authService) RefreshToken(params *dto_auth.RefreshTokenReq) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_auth.RefreshTokenReq)
		reply *pb_auth.RefreshTokenResp
	)
	req.RefreshToken = params.RefreshToken
	reply = s.authClient.RefreshToken(req)
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
	resp.Data = reply.AccessToken
	return
}
