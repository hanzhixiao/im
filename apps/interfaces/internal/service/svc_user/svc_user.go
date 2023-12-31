package svc_user

import (
	"im/apps/interfaces/internal/config"
	"im/apps/interfaces/internal/dto/dto_user"
	user_client "im/apps/user/client"
	"im/pkg/xhttp"
)

type UserService interface {
	UserList(params *dto_user.UserListReq) (resp *xhttp.Resp)
	EditUserInfo(params *dto_user.EditUserInfoReq, uid int64) (resp *xhttp.Resp)
	Search(params *dto_user.SearchUserReq, uid int64) (resp *xhttp.Resp)
	GetUserInfo(params *dto_user.UserInfoReq, uid int64) (resp *xhttp.Resp)
}

type userService struct {
	userClient user_client.UserClient
}

func NewUserService() UserService {
	conf := config.GetConfig()
	userClient := user_client.NewUserClient(conf.Etcd, conf.UserServer, conf.Jaeger, conf.Name)
	return &userService{userClient: userClient}
}
