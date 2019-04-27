package mq
import(
	"log"
)
type RabbitMq struct {
}
func (*RabbitMq)Push() bool{
	log.Printf("RabbitMq push success")
	return true
}
func (*RabbitMq)Consume() bool{
	log.Printf("RabbitMq consume success")
	return true
}