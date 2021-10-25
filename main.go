package main

import (
	"server/mq"
	"server/rt"
)

func main() {
	mq.InitDB()
	rt.Initserve()
}
