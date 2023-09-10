package router

import (
	"github.com/gin-gonic/gin"
	"im/apps/interfaces/dig"
	"im/apps/interfaces/internal/ctrl/ctrl_user"
	"im/apps/interfaces/internal/service/svc_user"
)

func registerUserRouter(group *gin.RouterGroup) {
	var svc svc_user.UserService
	dig.Invoke(func(s svc_user.UserService) {
		svc = s
	})
	ctrl := ctrl_user.NewUserCtrl(svc)
	router := group.Group("user")
	router.GET("list", ctrl.UserList)
	router.POST("edit_info", ctrl.EditUserInfo)
	router.GET("search", ctrl.Search)
	router.GET("user_info", ctrl.GetUserInfo)
}
