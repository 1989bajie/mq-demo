package mq
type KafkaMq struct {
}
func (*KafkaMq)Push() bool{
	return true
}
func (*KafkaMq)Consume() bool{
	return true
}