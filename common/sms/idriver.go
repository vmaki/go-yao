package sms

type IDriver interface {
	Send(phone string, message Message, config map[string]string) bool
}
