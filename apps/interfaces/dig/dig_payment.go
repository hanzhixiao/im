package dig

import (
	"im/apps/interfaces/internal/service/svc_payment"
)

func providePayment() {
	Provide(svc_payment.NewPaymentService)
}
