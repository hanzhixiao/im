package ctrl_convo

import (
	"im/apps/interfaces/internal/service/svc_convo"
)

type ConvoCtrl struct {
	convoService svc_convo.ConvoService
}

func NewConvoCtrl(convoService svc_convo.ConvoService) *ConvoCtrl {
	return &ConvoCtrl{convoService: convoService}
}
