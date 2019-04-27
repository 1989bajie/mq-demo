package main

import (
	"log"
	"mq-demo/mq"
	"mq-demo/conf"
	"strconv"
)
func main() {
	//mq测试
	//var mqObject 
	mqObject := mq.CreateMq(conf.QUEUE_TYPE_RABBIT)
	result := mqObject.Push()
	log.Printf("mq result:" + strconv.FormatBool(result))
}
