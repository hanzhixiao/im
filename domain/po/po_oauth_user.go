package po

import "im/pkg/entity"

type OauthUser struct {
	entity.GormEntityTs
	OauthId      int64  `gorm:"column:oauth_id;primary_key" json:"oauth_id"`        // 唯一ID
	Channel      int    `gorm:"column:channel;default:0;NOT NULL" json:"channel"`   // 1:Github 2:Google
	Openid       string `gorm:"column:openid;NOT NULL" json:"openid"`               // 第三方用户ID
	Uid          int64  `gorm:"column:uid;default:0;NOT NULL" json:"uid"`           // im uid
	Username     string `gorm:"column:username;NOT NULL" json:"username"`           // 第三方username
	Nickname     string `gorm:"column:nickname;NOT NULL" json:"nickname"`           // 第三方nickname
	Email        string `gorm:"column:email;NOT NULL" json:"email"`                 // Email
	AccessToken  string `gorm:"column:access_token;NOT NULL" json:"access_token"`   // 第三方AccessToken
	RefreshToken string `gorm:"column:refresh_token;NOT NULL" json:"refresh_token"` // 第三方RefreshToke
	Expire       int    `gorm:"column:expire;default:0;NOT NULL" json:"expire"`     // 过期时间 时间戳
	AvatarUrl    string `gorm:"column:avatar_url;NOT NULL" json:"avatar_url"`       // 第三方头像url
	Scope        string `gorm:"column:scope;NOT NULL" json:"scope"`                 // 用户授权的作用域，使用逗号（,）分隔
}
