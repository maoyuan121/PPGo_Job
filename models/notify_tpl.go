/************************************************************
** @Description: 通知模板
** @Author: Bee
** @Date:   2019-02-15 20:21
** @Last Modified by:   Bee
** @Last Modified time: 2019-02-15 20:21
*************************************************************/
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

// 模板类型
const (
	NotifyTplTypeSystem  = "system"
	NotifyTplTypeDefault = "default"
)

// 通知模板 Model
type NotifyTpl struct {
	Id         int    // 主键
	Type       string // 类型（system, default）
	TplName    string // 名称
	TplType    int    // 类型 （0：邮箱， 1：短信，2：钉钉，3：微信）
	Title      string // 标题
	Content    string // 内容
	Status     int
	CreateId   int   // 创建人
	UpdateId   int   // 更新人
	CreateTime int64 // 创建时间
	UpdateTime int64 // 更新时间
}

// 表名
func (t *NotifyTpl) TableName() string {
	return TableName("notify_tpl")
}

// 更新
func (t *NotifyTpl) Update(fields ...string) error {
	if t.TplName == "" {
		return fmt.Errorf("模板名称不能为空")
	}

	if t.Content == "" {
		return fmt.Errorf("模板内容不能为空")
	}

	if t.CreateTime == 0 {
		t.CreateTime = time.Now().Unix()
	}

	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

// 添加
func NotifyTplAdd(obj *NotifyTpl) (int64, error) {
	if obj.TplName == "" {
		return 0, fmt.Errorf("模板名称不能为空")
	}
	if obj.Content == "" {
		return 0, fmt.Errorf("模板内容不能为空")
	}
	if obj.CreateTime == 0 {
		obj.CreateTime = time.Now().Unix()
	}
	return orm.NewOrm().Insert(obj)
}

// 根据类型获取一个模板
func NotifyTplGetByTplType(tpl_type int, typestr string) (NotifyTpl, error) {
	var obj NotifyTpl
	err := orm.NewOrm().QueryTable(TableName("notify_tpl")).Filter("type", typestr).Filter("tpl_type", tpl_type).Filter("status", 1).Limit(1).One(&obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}

// 根据 ID 获取
func NotifyTplGetById(id int) (*NotifyTpl, error) {
	obj := &NotifyTpl{
		Id: id,
	}
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func NotifyTplGetByTplTypeList(tpl_type int) ([]*NotifyTpl, int64, error) {
	list := make([]*NotifyTpl, 0)
	total, err := orm.NewOrm().QueryTable(TableName("notify_tpl")).Filter("tpl_type", tpl_type).Filter("status", 1).All(&list)
	return list, total, err
}

// 删除
func NotifyTplDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("notify_tpl")).Filter("id", id).Delete()
	return err
}

// 查询
func NotifyTplGetList(page, pageSize int, filters ...interface{}) ([]*NotifyTpl, int64) {

	offset := (page - 1) * pageSize
	list := make([]*NotifyTpl, 0)
	query := orm.NewOrm().QueryTable(TableName("notify_tpl"))
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
