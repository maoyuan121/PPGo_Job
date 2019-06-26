/************************************************************
** @Description: 服务器
** @Author: haodaquan
** @Date:   2018-06-09 16:11
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-06-09 16:11
*************************************************************/
package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

// 服务器 Model
type TaskServer struct {
	Id             int    // 主键
	GroupId        int    // 所属分组
	ConnectionType int    // 连接类型（0：ssh，1：telnet）
	ServerName     string // 名称
	ServerAccount  string // 登录账号
	ServerOuterIp  string // 外网 IP
	ServerIp       string // 内网 IP
	Port           int    // 端口
	Password       string // 登录密码
	PrivateKeySrc  string // 私钥地址
	PublicKeySrc   string // 公钥地址
	Type           int    // 登录类型（0：密码，1：密钥）
	Detail         string // 备注
	CreateTime     int64  // 创建时间
	UpdateTime     int64  // 更新时间
	Status         int    // 状态（0：正常，1：禁用）
}

// 表名
func (t *TaskServer) TableName() string {
	return TableName("task_server")
}

// 更新
func (t *TaskServer) Update(fields ...string) error {
	if t.ServerName == "" {
		return fmt.Errorf("服务器名不能为空")
	}
	if t.ServerIp == "" {
		return fmt.Errorf("服务器IP不能为空")
	}

	if t.ServerAccount == "" {
		return fmt.Errorf("登录账户不能为空")
	}

	if t.Type == 0 && t.Password == "" {
		return fmt.Errorf("服务器密码不能为空")
	}

	if t.Type == 1 && t.PrivateKeySrc == "" {
		return fmt.Errorf("私钥不能为空")
	}

	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

// 新增
func TaskServerAdd(obj *TaskServer) (int64, error) {
	if obj.ServerName == "" {
		return 0, fmt.Errorf("服务器名不能为空")
	}
	if obj.ServerIp == "" {
		return 0, fmt.Errorf("服务器IP不能为空")
	}

	if obj.ServerAccount == "" {
		return 0, fmt.Errorf("登录账户不能为空")
	}

	if obj.Type == 0 && obj.Password == "" {
		return 0, fmt.Errorf("服务器密码不能为空")
	}

	if obj.Type == 1 && obj.PrivateKeySrc == "" {
		return 0, fmt.Errorf("私钥不能为空")
	}
	return orm.NewOrm().Insert(obj)
}

// 根据 ID 获取
func TaskServerGetById(id int) (*TaskServer, error) {
	obj := &TaskServer{
		Id: id,
	}
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

// 根据 ID 集合获取
func TaskServerGetByIds(ids string) ([]*TaskServer, int64) {

	serverFilters := make([]interface{}, 0)
	//serverFilters = append(serverFilters, "status", 1)

	TaskServerIdsArr := strings.Split(ids, ",")
	TaskServerIds := make([]int, 0)
	for _, v := range TaskServerIdsArr {
		id, _ := strconv.Atoi(v)
		TaskServerIds = append(TaskServerIds, id)
	}
	serverFilters = append(serverFilters, "id__in", TaskServerIds)
	return TaskServerGetList(1, 1000, serverFilters...)
}

// 删除
func TaskServerDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("task_server")).Filter("id", id).Delete()
	return err
}

// 查询
func TaskServerGetList(page, pageSize int, filters ...interface{}) ([]*TaskServer, int64) {

	offset := (page - 1) * pageSize
	list := make([]*TaskServer, 0)
	query := orm.NewOrm().QueryTable(TableName("task_server"))
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
