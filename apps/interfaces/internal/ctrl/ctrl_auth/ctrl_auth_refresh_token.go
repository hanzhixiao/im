package ctrl_auth

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/internal/dto/dto_auth"
	"im/pkg/common/xgin"
	"im/pkg/common/xlog"
	"im/pkg/xhttp"
)

func (ctrl *AuthCtrl) RefreshToken(ctx *gin.Context) {
	var (
		params = new(dto_auth.RefreshTokenReq)
		resp   *xhttp.Resp
		err    error
	)
	if err = xgin.BindJSON(ctx, params); err != nil {
		xlog.Warn(xhttp.ERROR_CODE_HTTP_REQ_PARAM_ERR, xhttp.ERROR_HTTP_REQ_PARAM_ERR, err.Error())
		return
	}
	resp = ctrl.authService.RefreshToken(params)
	if resp.Code > 0 {
		xhttp.Error(ctx, resp.Code, resp.Msg)
		return
	}
	xhttp.Success(ctx, resp.Data)
}
