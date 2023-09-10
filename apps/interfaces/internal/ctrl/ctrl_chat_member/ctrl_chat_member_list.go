package ctrl_chat_member

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/internal/dto/dto_chat_member"
	"im/pkg/common/xgin"
	"im/pkg/common/xlog"
	"im/pkg/xhttp"
)

func (ctrl *ChatMemberCtrl) ChatMemberList(ctx *gin.Context) {
	var (
		params = new(dto_chat_member.ChatMemberListReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.chatMemberService.ChatMemberList(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
