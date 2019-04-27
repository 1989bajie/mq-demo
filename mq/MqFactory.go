package mq
import(
	"mq-demo/conf"
)
func CreateMq(myType string) MyInterface{
	//简单工厂模式
	switch(myType) {
	case conf.QUEUE_TYPE_RABBIT:
		return &RabbitMq{}
	}
	return &RabbitMq{}
}