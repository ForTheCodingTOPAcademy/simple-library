package notifications

import (
	"fmt"
)

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
	EmailAdress string
}

type SMSNotifier struct {
	PhoneNumber string
}

func (e EmailNotifier) Notify(message string) {
	fmt.Printf("Отправляю email на %s: '%s'\n", e.EmailAdress, message)
}

func (s SMSNotifier) Notify(message string) {
	fmt.Printf("Отправляю SMS на номер %s: '%s'\n", s.PhoneNumber, message)
}
