package ctrl_chat

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/internal/dto/dto_chat"
	"im/pkg/common/xgin"
	"im/pkg/common/xlog"
	"im/pkg/xhttp"
)

func (ctrl *ChatCtrl) GroupChatDetails(ctx *gin.Context) {
	var (
		params = new(dto_chat.GroupChatDetailsReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.chatService.GroupChatDetails(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
