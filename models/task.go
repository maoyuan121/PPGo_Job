/************************************************************
** @Description: 任务
** @Author: haodaquan
** @Date:   2018-06-11 21:26
** @Last Modified by:   Bee
** @Last Modified time: 2019-02-15 21:32
*************************************************************/
package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

const (
	TASK_SUCCESS = 0  // 任务执行成功
	TASK_ERROR   = -1 // 任务执行出错
	TASK_TIMEOUT = -2 // 任务执行超时
)

// 任务 Model
type Task struct {
	Id            int    // 主键
	GroupId       int    // 所属分组
	ServerIds     string // 服务器集合
	TaskName      string // 名称
	Description   string // 备注
	CronSpec      string // 时间表达式
	Concurrent    int    // 是否并发（0：单例，1：并发）
	Command       string // 命令脚本
	Timeout       int    // 超时
	ExecuteTimes  int    // 执行次数
	PrevTime      int64  // 上次执行时间
	Status        int    // 状态  （0：成功，-1：出错，-2：超时）
	IsNotify      int    // 出错的时候发通知（1：是，0：否）
	NotifyType    int    // 通知类型（0：邮件，1：短信，2：钉钉，3：微信）
	NotifyTplId   int    // 通知模板
	NotifyUserIds string // 通知用户
	CreateId      int    // 创建人
	UpdateId      int    // 更新人
	CreateTime    int64  // 创建时间
	UpdateTime    int64  // 更新时间
}

// 表名
func (t *Task) TableName() string {
	return TableName("task")
}

// 更新
func (t *Task) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

// 新增
func TaskAdd(task *Task) (int64, error) {
	if task.TaskName == "" {
		return 0, fmt.Errorf("任务名称不能为空")
	}

	if task.CronSpec == "" {
		return 0, fmt.Errorf("时间表达式不能为空")
	}
	if task.Command == "" {
		return 0, fmt.Errorf("命令内容不能为空")
	}
	if task.CreateTime == 0 {
		task.CreateTime = time.Now().Unix()
	}
	return orm.NewOrm().Insert(task)
}

// 查询
func TaskGetList(page, pageSize int, filters ...interface{}) ([]*Task, int64) {
	offset := (page - 1) * pageSize

	tasks := make([]*Task, 0)

	query := orm.NewOrm().QueryTable(TableName("task"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-status", "task_name", "-id").Limit(pageSize, offset).All(&tasks)

	return tasks, total
}

// 根据分组获取
func TaskResetGroupId(groupId int) (int64, error) {
	return orm.NewOrm().QueryTable(TableName("task")).Filter("group_id", groupId).Update(orm.Params{
		"group_id": 0,
	})
}

// 根据 ID 获取
func TaskGetById(id int) (*Task, error) {
	task := &Task{
		Id: id,
	}

	err := orm.NewOrm().Read(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// 删除（逻辑删除）
func TaskDel(id int) (int64, error) {
	return orm.NewOrm().QueryTable(TableName("task")).Filter("id", id).Update(orm.Params{
		"status": -1,
	})
	//_, err := orm.NewOrm().QueryTable(TableName("task")).Filter("id", id).Delete()
	//return err
}
