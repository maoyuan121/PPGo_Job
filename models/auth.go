/**********************************************
** @Des: 权限 Model
** @Author: haodaquan
** @Date:   2017-09-09 20:50:36
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 21:42:08
***********************************************/
package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// 权限模型
type Auth struct {
	Id         int    // 主键
	AuthName   string // 权限名
	AuthUrl    string // 菜单地址
	UserId     int    // 最后创建或更新人
	Pid        int    // 上级权限 ID
	Sort       int    // 排序
	Icon       string // 图标字体
	IsShow     int    // 是否左侧导航栏显示
	Status     int    // 状态 0 删除， 1 正常
	CreateId   int    // 创建人
	UpdateId   int    // 更新人
	CreateTime int64  // 创建时间
	UpdateTime int64  // 更新时间
}

// 表名
func (a *Auth) TableName() string {
	return TableName("uc_auth")
}

// 权限列表
func AuthGetList(page, pageSize int, filters ...interface{}) ([]*Auth, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Auth, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_auth"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("pid", "sort").Limit(pageSize, offset).All(&list)

	return list, total
}

// 如果是超级管理员查所有的权限， 否则根据 authIds 查找
func AuthGetListByIds(authIds string, userId int) ([]*Auth, error) {

	list1 := make([]*Auth, 0)
	var list []orm.Params
	//list:=[]orm.Params
	var err error
	if userId == 1 {
		//超级管理员
		_, err = orm.NewOrm().Raw("select id,auth_name,auth_url,pid,icon,is_show from pp_uc_auth where status=? order by pid asc,sort asc", 1).Values(&list)
	} else {
		_, err = orm.NewOrm().Raw("select id,auth_name,auth_url,pid,icon,is_show from pp_uc_auth where status=1 and id in("+authIds+") order by pid asc,sort asc", authIds).Values(&list)
	}

	for k, v := range list {
		fmt.Println(k, v)
	}

	fmt.Println(list)
	return list1, err
}

// 创建权限
func AuthAdd(auth *Auth) (int64, error) {
	return orm.NewOrm().Insert(auth)
}

// 根据 ID 查找权限
func AuthGetById(id int) (*Auth, error) {
	a := new(Auth)

	err := orm.NewOrm().QueryTable(TableName("uc_auth")).Filter("id", id).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// 更新权限
func (a *Auth) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
