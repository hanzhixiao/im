package ctrl_chat_member

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/internal/dto/dto_chat_member"
	"im/pkg/common/xgin"
	"im/pkg/common/xlog"
	"im/pkg/xhttp"
)

func (ctrl *ChatMemberCtrl) ContactList(ctx *gin.Context) {
	var (
		params = new(dto_chat_member.ContactListReq)
		uid    int64
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.ShouldBindQuery(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	uid = xgin.GetUid(ctx)
	if uid == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	resp = ctrl.chatMemberService.ContactList(params, uid)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
