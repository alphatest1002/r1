package models

import (
	//导入mysql的驱动程序init函数
	_ "github.com/go-sql-driver/mysql"
)

//封装表,映射（orm模型）: object replationship mapping 对象关系映射模型
type Students struct {
	Id        int64  `db:"id"`
	Name      string `db:"name"`
	Age       int    `db:"age"`
	ClassId   int64  `db:"classid"`
	ClassName string `db:"classname"`
}
