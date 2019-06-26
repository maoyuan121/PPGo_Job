/************************************************************
** @Description: PPGo_Job2
** @Author: haodaquan
** @Date:   2018-06-05 22:24
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-06-05 22:24
*************************************************************/
package main

import (
	"github.com/astaxie/beego"
	"PPGo_Job/jobs"
	"PPGo_Job/models"
	_ "PPGo_Job/routers"
	"time"
)

func init() {
	//初始化数据模型
	var StartTime = time.Now().Unix()
	models.Init(StartTime)
	jobs.InitJobs()
}

func main() {
	beego.Run()
}
