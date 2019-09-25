package controller

import (
	"database/sql"
	"fdStudent/models"

	"log"
	"net/http"
	"text/template"
	//导入mysql的驱动程序init函数
	_ "github.com/go-sql-driver/mysql"
)

type StudentController struct {
	DB *sql.DB //此属性为了关联models
}

//获取所有学生列表的方法
func (stu *StudentController) FetchAll() []models.Students {
	//声明Students的orm模型
	student := models.Students{}
	//声明一个装载多个学生的切片容器
	var stuData []models.Students

	rows, err := stu.DB.Query("SELECT stu.id AS id,`name`,`age`,`classid`,`classname` FROM students AS stu LEFT JOIN classes AS cls ON stu.classid=cls.id")
	if err != nil {
		log.Panic("查询学生列表有误", err)
	}
	for rows.Next() {
		//绑定一次记录到orm模型当中
		rows.Scan(&student.Id, &student.Name, &student.Age, &student.ClassId, &student.ClassName)
		//把结果放到切片中
		stuData = append(stuData, student)
	}
	return stuData
}

//实现ServeHTTP的接口(为了分发)
func (stu *StudentController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//传递模板的数据
	tmp := make(map[string]interface{})
	//获取路由的URI
	uri := r.URL.Path
	switch uri {
	case "/":
		//获取所有的列表放置到stuData里面交给模板
		tmp["stuData"] = stu.FetchAll()
		//解析首页模板
		tpl, _ := template.ParseFiles("./view/students/index.html")
		tpl.Execute(w, tmp)
	}
}
