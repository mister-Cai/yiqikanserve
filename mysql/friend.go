package mysql

import (
	"fmt"
	"server/mq"
)

func init() {
	// friend := Selectfriend(4)
	// // friend := Selectfriendall(1)
	// friend := Upfriend(4, "123465798132")
	// fmt.Println(friend)
	// friend := Delfriend(4)
	// fmt.Println(friend)
	// friend := Friend{
	// 	UserName:   "111",
	// 	Password:   "123456798",
	// 	CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	// }
	// friend2 := Addfriend(friend)
	// fmt.Println(friend2)
}

type Friend struct {
	Id  int
	Uid int
	Fid string
}

func Addfriends(friend Friend) bool {
	//开启事务
	tx, err := mq.DB.Begin()
	if err != nil {
		fmt.Println("Addfriend Begin fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO friend (`id`,`uid`) VALUES (?,?)")
	if err != nil {
		fmt.Println("Inster Prepare fail")
		return false
	}
	res, err := stmt.Exec(friend.Id, friend.Uid)
	if err != nil {
		fmt.Println("inster fail", err)
		return false
	}
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func Delfriend(Id int) bool {
	//开启事务
	tx, err := mq.DB.Begin()
	if err != nil {
		fmt.Println("Delfriend Begin fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM friend WHERE id=?")
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
func Upfriend(uid int, fid int) bool {
	//开启事务
	tx, err := mq.DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE friend SET fid=? WHERE id=?")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}

	//放入参数并执行sql语句
	res, err := stmt.Exec(fid, uid)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func Selectfriendall(Id int) []Friend {
	var friends []Friend
	rows, err := mq.DB.Query("SELECT * FROM friend WHERE id>?", Id)
	if err != nil {
		fmt.Println("查询出错了", err)
	}
	defer rows.Close()
	for rows.Next() {
		var friend Friend
		err := rows.Scan(&friend.Id, &friend.Uid, &friend.Fid)
		if err != nil {
			fmt.Println("查询出错了", err)
		}
		friends = append(friends, friend)
	}
	return friends
}
func Selectfriend(Id int) Friend {
	var friend Friend
	err := mq.DB.QueryRow("SELECT * FROM friend WHERE id=?", Id).Scan(&friend.Id, &friend.Uid, &friend.Fid)
	if err != nil {
		fmt.Println("查询出错了", err)
	}
	return friend
}
