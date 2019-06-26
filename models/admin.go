/**********************************************
** @Des: 管理员 Model
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

// 管理员模型
type Admin struct {
	Id         int    // 主键
	LoginName  string // 用户名
	RealName   string // 真实姓名
	Password   string // 密码
	RoleIds    string // 角色集合（格式：1，2，3，4）
	Phone      string // 电话
	Email      string // Email
	Dingtalk   string // 钉钉
	Wechat     string // 微信
	Salt       string // 密码盐
	LastLogin  int64  // 最后一次登录时间
	LastIp     string // 最后一次登录 IP 地址
	Status     int    // 状态 （0：禁用，1：正常）
	CreateId   int    // 创建者
	UpdateId   int    // 更新者
	CreateTime int64  // 创建时间
	UpdateTime int64  // 更新时间
}

// 表名
func (a *Admin) TableName() string {
	return TableName("uc_admin")
}

// 新增管理员
func AdminAdd(a *Admin) (int64, error) {
	return orm.NewOrm().Insert(a)
}

// 根据登录名获取管理员
func AdminGetByName(loginName string) (*Admin, error) {
	a := new(Admin)
	err := orm.NewOrm().QueryTable(TableName("uc_admin")).Filter("login_name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 管理员列表
func AdminGetList(page, pageSize int, filters ...interface{}) ([]*Admin, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Admin, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_admin"))
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

// 根据 ID 查询管理员
func AdminGetById(id int) (*Admin, error) {
	r := new(Admin)
	err := orm.NewOrm().QueryTable(TableName("uc_admin")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// 更新管理员
func (a *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
