package ctrl_user

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/internal/dto/dto_user"
	"im/pkg/common/xgin"
	"im/pkg/common/xlog"
	"im/pkg/xhttp"
)

func (ctrl *UserCtrl) EditUserInfo(ctx *gin.Context) {
	var (
		params = new(dto_user.EditUserInfoReq)
		uid    int64
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.BindJSON(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	uid = xgin.GetUid(ctx)
	if uid == 0 {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_USER_ID_DOESNOT_EXIST, xhttp.ERROR_HTTP_USER_ID_DOESNOT_EXIST)
		return
	}
	resp = ctrl.userService.EditUserInfo(params, uid)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
