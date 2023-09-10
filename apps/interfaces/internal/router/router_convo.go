package router

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/dig"
	"im/apps/interfaces/internal/ctrl/ctrl_convo"
	"im/apps/interfaces/internal/service/svc_convo"
)

func registerConvoRouter(group *gin.RouterGroup) {
	var svc svc_convo.ConvoService
	dig.Invoke(func(s svc_convo.ConvoService) {
		svc = s
	})
	ctrl := ctrl_convo.NewConvoCtrl(svc)
	router := group.Group("convo")
	router.POST("list", ctrl.ConvoList)
	router.POST("chat_seq_list", ctrl.ConvoChatSeqList)
}
