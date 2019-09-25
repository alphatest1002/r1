package models

import (
	"fmt"
	"log"
	//导入数据库配置
	"fdStudent/common"
	//导入数据库包
	"database/sql"
	//导入mysql的驱动程序init函数
	_ "github.com/go-sql-driver/mysql"
)

func GetDataBase() *sql.DB {
	//配置dns
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", common.DB_USER, common.DB_PASSWORD, common.DB_HOST, common.DB_PORT, common.DB_NAME)
	//打开连接
	db, err := sql.Open(common.DB_DRIVER, dns)
	if err != nil {
		log.Panic("连接数据库失败了", err)
	}
	return db

}
