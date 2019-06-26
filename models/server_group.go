/************************************************************
** @Description: 服务器分组
** @Author: haodaquan
** @Date:   2018-06-08 21:49
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-06-08 21:49
*************************************************************/
package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// 服务器分组名
type ServerGroup struct {
	Id          int    // 主键
	CreateId    int    // 创建人
	UpdateId    int    // 更新人
	GroupName   string // 名称
	Description string // 备注
	CreateTime  int64  // 创建时间
	UpdateTime  int64  // 更新时间
	Status      int    // 状态 （0：删除，1：正常）
}

// 表名
func (t *ServerGroup) TableName() string {
	return TableName("task_server_group")
}

// 更新
func (t *ServerGroup) Update(fields ...string) error {
	if t.GroupName == "" {
		return fmt.Errorf("组名不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

// 创建
func ServerGroupAdd(obj *ServerGroup) (int64, error) {
	if obj.GroupName == "" {
		return 0, fmt.Errorf("组名不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

// 根据 ID 获取
func TaskGroupGetById(id int) (*ServerGroup, error) {
	obj := &ServerGroup{
		Id: id,
	}
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// 删除
func ServerGroupDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("task_server_group")).Filter("id", id).Delete()
	return err
}

// 查询
func ServerGroupGetList(page, pageSize int, filters ...interface{}) ([]*ServerGroup, int64) {
	offset := (page - 1) * pageSize
	list := make([]*ServerGroup, 0)
	query := orm.NewOrm().QueryTable(TableName("task_server_group"))
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
