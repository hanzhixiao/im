package ctrl_payment

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/internal/service/svc_payment"
	"im/pkg/xhttp"
)

type PaymentCtrl struct {
	paymentService svc_payment.PaymentService
}

func NewPaymentCtrl(paymentService svc_payment.PaymentService) *PaymentCtrl {
	return &PaymentCtrl{paymentService: paymentService}
}

func (ctrl *PaymentCtrl) Edit(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}

func (ctrl *PaymentCtrl) Info(ctx *gin.Context) {
	xhttp.Success(ctx, nil)
}
