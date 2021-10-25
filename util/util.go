package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//获取指定长度的随机数字
func RandNum(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	fmt.Println(sb.String())
	return sb.String()
}
