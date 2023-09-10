package svc_red_env_receive

import (
	"github.com/jinzhu/copier"
	"im/apps/interfaces/internal/dto/dto_red_env_receive"
	"im/pkg/common/xlog"
	"im/pkg/proto/pb_red_env_receive"
	"im/pkg/xhttp"
)

func (s *redEnvReceiveService) GrabRedEnvelope(params *dto_red_env_receive.GrabRedEnvelopeReq, uid int64) (resp *xhttp.Resp) {
	resp = new(xhttp.Resp)
	var (
		req   = new(pb_red_env_receive.GrabRedEnvelopeReq)
		reply *pb_red_env_receive.GrabRedEnvelopeResp
	)
	copier.Copy(req, params)
	req.Uid = uid
	reply = s.redEnvReceiveClient.GrabRedEnvelope(req)
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
	resp.Data = reply.Result
	return
}
