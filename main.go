package main

import (
	"fdStudent/controller"
	"fdStudent/models"
	"net/http"
	"time"
)

func main() {
	//连接数据库
	DB := models.GetDataBase()
	//声明学生的控制器
	stuController := &controller.StudentController{DB}
	//声明规划的路由
	http.Handle("/", stuController)               //首页(显示所有学生列表)
	http.Handle("/student/add", stuController)    //添加页面
	http.Handle("/student/insert", stuController) //添加入库后台逻辑
	http.Handle("/student/del", stuController)    //删除表的后台逻辑
	http.Handle("/student/editor", stuController) //编辑页面
	http.Handle("/student/update", stuController) //编辑入库的逻辑

	//配置服务器
	app := http.Server{
		Addr:        ":8888",
		Handler:     stuController,
		ReadTimeout: 2 * time.Second,
	}

	//启动服务器
	app.ListenAndServe()

}
