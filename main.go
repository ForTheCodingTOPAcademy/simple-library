package main

import (
	"fmt"
)

func main() {
	r := Reader{ID: 10, Name: "Саня", IsActive: false}
	b := Book{ID: 1, Title: "Вокруг света за 80 дней", Author: "Жуль Верн", IsIssued: false, ReaderID: r.ID}
	fmt.Println(b.String())
	b.ReturnBook()
	b.IssueBook(r)
	r.IsActive = true
	b.IssueBook(r)
	fmt.Println(b.String())
	b.AssignBooks()
	b.ReturnBook()
	fmt.Println(b.String())

	var emailN Notifier = EmailNotifier{EmailAdress: "example@gmail.com"}
	var SMSN Notifier = SMSNotifier{PhoneNumber: "+7 988 376 88 30"}

	SMSN.Notify("Hello through SMS")
	emailN.Notify("Hello through email")
}
