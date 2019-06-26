/**********************************************
** @Des: 角色拥有哪些权限关联模型
** @Author: haodaquan
** @Date:   2017-09-15 11:44:13
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:49:13
***********************************************/
package models

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

// 角色拥有哪些权限关联模型
type RoleAuth struct {
	AuthId int   `orm:"pk"` // 权限主键
	RoleId int64 ``         // 角色主键
}

// 表名
func (ra *RoleAuth) TableName() string {
	return TableName("uc_role_auth")
}

// 创建
func RoleAuthAdd(ra *RoleAuth) (int64, error) {
	return orm.NewOrm().Insert(ra)
}

// 批量创建
func RoleAuthBatchAdd(ras *[]RoleAuth) (int64, error) {
	return orm.NewOrm().InsertMulti(100, ras)
}

// 根据角色获取其拥有的权限
func RoleAuthGetById(id int) ([]*RoleAuth, error) {
	list := make([]*RoleAuth, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_role_auth"))
	_, err := query.Filter("role_id", id).All(&list, "AuthId")
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 删除某个角色拥有的所有权限
func RoleAuthDelete(id int) (int64, error) {
	_, err := orm.NewOrm().Raw("DELETE FROM `pp_uc_role_auth` WHERE `role_id` = ?",
		strconv.Itoa(id)).Exec()
	return 0, err
}

// 获取这些角色拥有的所有的权限的 ID
func RoleAuthGetByIds(RoleIds string) (Authids string, err error) {
	list := make([]*RoleAuth, 0)
	query := orm.NewOrm().QueryTable(TableName("uc_role_auth"))
	ids := strings.Split(RoleIds, ",")
	_, err = query.Filter("role_id__in", ids).All(&list, "AuthId")
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range list {
		if v.AuthId != 0 && v.AuthId != 1 {
			b.WriteString(strconv.Itoa(v.AuthId))
			b.WriteString(",")
		}
	}
	Authids = strings.TrimRight(b.String(), ",")
	return Authids, nil
}

// 批量新增权限
func RoleAuthMultiAdd(ras []*RoleAuth) (n int, err error) {
	query := orm.NewOrm().QueryTable(TableName("uc_role_auth"))
	i, _ := query.PrepareInsert()
	for _, ra := range ras {
		_, err := i.Insert(ra)
		if err == nil {
			n = n + 1
		}
	}
	i.Close() // 别忘记关闭 statement
	return n, err
}
