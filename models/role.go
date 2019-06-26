/**********************************************
** @Des: 角色
** @Author: haodaquan
** @Date:   2017-09-14 15:24:51
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:48:52
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

// 角色 Model
type Role struct {
	Id             int    // 主键
	RoleName       string // 名称
	Detail         string // 备注
	ServerGroupIds string // 资源分组
	TaskGroupIds   string // 任务分组
	Status         int    // 状态 （0： 删除， 1：正常）
	CreateId       int    // 创建人
	UpdateId       int    // 更新人
	CreateTime     int64  // 创建时间
	UpdateTime     int64  // 更新时间
}

// 表名
func (a *Role) TableName() string {
	return TableName("uc_role")
}

// 查询
func RoleGetList(page, pageSize int, filters ...interface{}) ([]*Role, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Role, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_role"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

// 创建
func RoleAdd(role *Role) (int64, error) {
	id, err := orm.NewOrm().Insert(role)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 根据 ID 获取
func RoleGetById(id int) (*Role, error) {
	r := new(Role)
	err := orm.NewOrm().QueryTable(TableName("uc_role")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// 更新
func (r *Role) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
