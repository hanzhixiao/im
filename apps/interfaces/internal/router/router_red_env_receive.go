package router

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/dig"
	"im/apps/interfaces/internal/ctrl/ctrl_red_env_receive"
	"im/apps/interfaces/internal/service/svc_red_env_receive"
)

func registerRedEnvReceiveRouter(group *gin.RouterGroup) {
	var svc svc_red_env_receive.RedEnvReceiveService
	dig.Invoke(func(s svc_red_env_receive.RedEnvReceiveService) {
		svc = s
	})
	ctrl := ctrl_red_env_receive.NewRedEnvReceiveCtrl(svc)
	router := group.Group("red_env_receive")
	router.POST("grab", ctrl.GrabRedEnvelope)
	router.POST("open", ctrl.OpenRedEnvelope)
}
