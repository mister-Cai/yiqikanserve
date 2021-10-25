package mysql

import (
	"fmt"
	"server/mq"
	"time"

	"github.com/beinan/fastid"
)

func init() {
	// initDB()
	// user := Selectuser(4)
	// // user := Selectuserall(1)
	// user := Upuser(4, "123465798132")
	// fmt.Println(user)
	// user := Deluser(4)
	// fmt.Println(user)
	// user1 := User{
	// 	UserName:   "111",
	// 	Password:   "123456798",
	// 	CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	// }
	// user2 := Adduser(user1)
	// fmt.Println(user2)
}

type User struct {
	Id         int
	UserName   string
	Phone      string
	Password   string
	CreateTime string
	Power      int
}

func Adduser(user User) bool {
	//开启事务
	tx, err := mq.DB.Begin()
	if err != nil {
		fmt.Println("Adduser Begin fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO user (`id`,`phone`,`password`,`create_time`,`power`) VALUES (?,?,?,?,?)")
	if err != nil {
		fmt.Println("Inster Prepare fail")
		return false
	}
	res, err := stmt.Exec(fastid.CommonConfig.GenInt64ID(), user.Phone, user.Password, time.Now().Format("2006-01-02 15:04:05"), user.Power)
	if err != nil {
		fmt.Println("inster fail", err)
		return false
	}
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func Deluser(Id int) bool {
	//开启事务
	tx, err := mq.DB.Begin()
	if err != nil {
		fmt.Println("Deluser Begin fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		fmt.Println("Delete Prepare fail")
		return false
	}
	//放入参数并执行sql语句
	res, err := stmt.Exec(Id)
	if err != nil {
		fmt.Println("delete fail", err)
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func Upuser(Id int, phone string) bool {
	//开启事务
	tx, err := mq.DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE user SET phone=? WHERE id=?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//放入参数并执行sql语句
	res, err := stmt.Exec(phone, Id)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func Selectuserall(Id int) []User {
	var users []User
	rows, err := mq.DB.Query("SELECT * FROM user WHERE id>?", Id)
	if err != nil {
		fmt.Println("查询出错了", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.UserName, &user.Phone, &user.Password, &user.CreateTime, &user.Power)
		if err != nil {
			fmt.Println("查询出错了", err)
		}
		users = append(users, user)
	}
	return users
}
func Selectuserid(Id int) User {
	var user User
	err := mq.DB.QueryRow("SELECT * FROM user WHERE id=?", Id).Scan(&user.Id, &user.UserName, &user.Phone, &user.Password, &user.CreateTime, &user.Power)
	if err != nil {
		fmt.Println("查询出错了", err)
	}
	return user
}
func Selectuserphone(Phone string) User {
	var user User
	err := mq.DB.QueryRow("SELECT id, IFNULL(username,''),phone,password,create_time,power FROM user WHERE phone=?", Phone).Scan(&user.Id, &user.UserName, &user.Phone, &user.Password, &user.CreateTime, &user.Power)
	if err != nil {
		fmt.Println("查询出错了", err)
	}
	return user
}
