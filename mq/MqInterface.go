package mq
type MyInterface interface {
	Push() bool
	Consume() bool
}