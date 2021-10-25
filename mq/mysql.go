package mq

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	USERNAME = "root"
	PASSWORD = "123456"
	IP       = "47.103.114.136"
	PORT     = "3306"
	dbNAME   = "yiqikan"
)

func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", IP, ":", PORT, ")/", dbNAME, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path) //连接数据库
	DB.SetConnMaxLifetime(100)      //最大连接数
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println("connect fail")
		return
	}
	fmt.Println("connect success")
}
