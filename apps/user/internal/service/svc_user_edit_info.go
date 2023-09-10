package service

import (
	"context"
	"gorm.io/gorm"
	"im/pkg/common/xlog"
	"im/pkg/common/xmysql"
	"im/pkg/constant"
	"im/pkg/entity"
	"im/pkg/proto/pb_user"
)

func (s *userService) EditUserInfo(ctx context.Context, req *pb_user.EditUserInfoReq) (resp *pb_user.EditUserInfoResp, _ error) {
	resp = new(pb_user.EditUserInfoResp)
	var (
		u        = entity.NewMysqlUpdate()
		imId     interface{}
		mobile   interface{}
		nickname interface{}
		ok       bool
		err      error
	)

	defer func() {
		if err != nil {
			xlog.Warn(resp.Code, resp.Msg, err.Error())
		}
	}()

	u.SetFilter("uid=?", req.Uid)
	req.Kvs.StrFieldValidation(userUpdateFields, u.Values)
	req.Kvs.IntFieldValidation(userUpdateFields, u.Values)

	err = xmysql.Transaction(func(tx *gorm.DB) (err error) {
		// im ID 重复校验
		if imId, ok = u.Values["im_id"]; ok == true {
			switch imId.(type) {
			case string:
				err = s.RecheckimId(tx, req.Uid, imId.(string), resp)
				if err != nil {
					return
				}
			default:
				err = ERR_USER_PARAM_ERR
				resp.Set(ERROR_CODE_USER_PARAM_ERR, ERROR_USER_PARAM_ERR)
				return
			}
		}

		// mobile 重复校验
		if mobile, ok = u.Values["mobile"]; ok == true {
			switch mobile.(type) {
			case string:
				err = s.RecheckMobile(tx, req.Uid, mobile.(string), resp)
				if err != nil {
					return
				}
			default:
				err = ERR_USER_PARAM_ERR
				resp.Set(ERROR_CODE_USER_PARAM_ERR, ERROR_USER_PARAM_ERR)
				return
			}
		}

		err = s.userRepo.TxUpdateUser(tx, u)
		if err != nil {
			/*
				switch err.(type) {
				case *mysql.MySQLError:
					if err.(*mysql.MySQLError).Number == constant.ERROR_CODE_MYSQL_DUPLICATE_ENTRY {
						if strings.HasSuffix(err.(*mysql.MySQLError).Message, constant.DUPLICATE_ENTRY_KV_USERS_MOBILE) {
							resp.Set(ERROR_CODE_USER_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT, ERROR_USER_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT)
							return
						}
						if strings.HasSuffix(err.(*mysql.MySQLError).Message, constant.DUPLICATE_ENTRY_KV_USERS_im_ID) {
							resp.Set(ERROR_CODE_USER_im_ID_HAS_BEEN_OCCUPIED, ERROR_USER_im_ID_HAS_BEEN_OCCUPIED)
							return
						}
					}
				}
			*/
			resp.Set(ERROR_CODE_USER_UPDATE_VALUE_FAILED, ERROR_USER_UPDATE_VALUE_FAILED)
			return
		}

		if nickname, ok = u.Values["nickname"]; ok == false {
			return
		}
		u.Reset()
		u.SetFilter("uid=?", req.Uid)
		u.SetFilter("sync=?", constant.SYNCHRONIZE_USER_INFO)
		u.Set("alias", nickname)
		err = s.chatMemberRepo.TxUpdateChatMember(tx, u)
		if err != nil {
			resp.Set(ERROR_CODE_USER_UPDATE_VALUE_FAILED, ERROR_USER_UPDATE_VALUE_FAILED)
			return
		}
		return
	})
	if err != nil {
		return
	}
	// 删除缓存
	err = s.userCache.DelUserInfo(req.Uid)
	if err != nil {
		resp.Set(ERROR_CODE_USER_UPDATE_USER_CACHE_FAILED, ERROR_USER_UPDATE_USER_CACHE_FAILED)
		return
	}

	// 更新缓存
	go s.updateChatMemberCacheInfo(req.Uid)
	return
}

func (s *userService) RecheckimId(tx *gorm.DB, uid int64, imId string, resp *pb_user.EditUserInfoResp) (err error) {
	var (
		w      = entity.NewMysqlQuery()
		exists bool
	)
	w.SetFilter("im_id=?", imId)
	exists, err = s.userRepo.TxExists(tx, w, uid)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		return
	}
	if exists == true {
		err = ERR_USER_im_ID_HAS_BEEN_OCCUPIED
		resp.Set(ERROR_CODE_USER_im_ID_HAS_BEEN_OCCUPIED, ERROR_USER_im_ID_HAS_BEEN_OCCUPIED)
		return
	}
	return
}

func (s *userService) RecheckMobile(tx *gorm.DB, uid int64, mobile string, resp *pb_user.EditUserInfoResp) (err error) {
	var (
		w      = entity.NewMysqlQuery()
		exists bool
	)
	w.SetFilter("mobile=?", mobile)
	exists, err = s.userRepo.TxExists(tx, w, uid)
	if err != nil {
		resp.Set(ERROR_CODE_USER_QUERY_DB_FAILED, ERROR_USER_QUERY_DB_FAILED)
		return
	}
	if exists == true {
		err = ERR_USER_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT
		resp.Set(ERROR_CODE_USER_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT, ERROR_USER_THE_MOBILE_HAS_BEEN_BOUND_TO_ANOTHER_ACCOUNT)
		return
	}
	return
}
