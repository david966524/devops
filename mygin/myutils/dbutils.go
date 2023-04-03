package myutils

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("run dbuitls")
}

// 基于database/sql连接，进行二次封装
func ConnectMysqlByDatabaseSql() (*gorm.DB, error) {
	config := GetConfig()
	host := config.Get("mysql.ip")
	port := config.Get("mysql.port")
	user := config.Get("mysql.user")
	pass := config.Get("mysql.password")
	dbname := config.Get("mysql.database")

	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(user)
	fmt.Println(pass)
	fmt.Println(dbname)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)

	fmt.Println("当前jdbc地址:    " + dns)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	return gorm.Open(mysql.New(mysql.Config{Conn: db}))
}
