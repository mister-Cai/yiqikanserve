package mysql

import (
	"fmt"
	"server/mq"
)

var id int

func Selectinvit(Invitcode string) bool {
	err := mq.DB.QueryRow("SELECT id FROM invitCode WHERE code=?", Invitcode).Scan(&id)
	fmt.Println(id)
	if err != nil {
		fmt.Println("查询出错了", err)
		return false
	}
	return true
}
