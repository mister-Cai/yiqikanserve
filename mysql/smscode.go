package mysql

import (
	"fmt"
	"server/mq"
)

func init() {
}

type Smscode struct {
	Id      int
	Phone   string
	Code    string
	Endtime int64
}

func Addsms(sms Smscode) bool {
	//开启事务
	tx, err := mq.DB.Begin()
	if err != nil {
		fmt.Println("Adduser Begin fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO smscode (`phone`,`code`,endtime) VALUES (?,?,?)")
	if err != nil {
		fmt.Println("Inster Prepare fail")
		return false
	}
	res, err := stmt.Exec(sms.Phone, sms.Code, sms.Endtime)
	if err != nil {
		fmt.Println("inster fail", err)
		return false
	}
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func Delsms(Code string) bool {
	//开启事务
	tx, err := mq.DB.Begin()
	if err != nil {
		fmt.Println("Deluser Begin fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM smscode WHERE code=?")
	if err != nil {
		fmt.Println("Delete Prepare fail")
		return false
	}
	//放入参数并执行sql语句
	res, err := stmt.Exec(Code)
	if err != nil {
		fmt.Println("delete fail", err)
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func Selectsms(Code string) Smscode {
	var sms Smscode
	err := mq.DB.QueryRow("SELECT * FROM smscode WHERE code=?", Code).Scan(&sms.Id, &sms.Phone, &sms.Code, &sms.Endtime)
	if err != nil {
		fmt.Println("查询出错了", err)
		return sms
	}
	return sms
}
