package do

import (
	"im/domain/po"
	"im/pkg/proto/pb_auth"
)

type SignIn struct {
	User         *po.User
	Avatar       *po.Avatar
	AccessToken  *pb_auth.Token
	RefreshToken *pb_auth.Token
	Code         int32
	Msg          string
	Err          error
}
