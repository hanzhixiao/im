package repo

import (
	"im/domain/po"
	"im/pkg/common/xmysql"
	"im/pkg/entity"
)

func (r *userRepository) QueryUser(q *entity.MysqlQuery, dist interface{}) (err error) {
	db := xmysql.GetDB()
	q.Model = new(po.User)
	err = q.Find(db, dist)
	//err = db.Model(po.User{}).Select(q.Fields).Where(q.Query, q.Args...).Find(dist).Error
	return
}
