package router

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/dig"
	"im/apps/interfaces/internal/ctrl/ctrl_payment"
	"im/apps/interfaces/internal/service/svc_payment"
)

func registerPaymentRouter(group *gin.RouterGroup) {
	var svc svc_payment.PaymentService
	dig.Invoke(func(s svc_payment.PaymentService) {
		svc = s
	})
	ctrl := ctrl_payment.NewPaymentCtrl(svc)
	router := group.Group("payment")
	router.POST("edit", ctrl.Edit)
	router.GET("info", ctrl.Info)
}
