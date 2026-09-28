package main

import (
	"fmt"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	IsIssued bool
	ReaderID int
}

type Reader struct {
	ID       int
	Name     string
	IsActive bool
}

func (b Book) String() string {
	id := 0
	if b.IsIssued {
		id = b.ReaderID
	}
	return fmt.Sprintf("-----------------------------------\n%d: %s by %s. Выдана: %v. ID пользователя: %d\n-----------------------------------", b.ID, b.Title, b.Author, b.IsIssued, id)
}

func (b *Book) IssueBook(r Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s by %s уже кому-то выдана читателю \n", b.Title, b.Author)
		return
	}
	if !r.IsActive {
		fmt.Printf("Читатель %s неактивен и не может получить книгу\n", r.Name)
		return
	}
	b.ReaderID = r.ID
	b.IsIssued = true
	fmt.Printf("Книга %s выдана %s\n", b.Title, r.Name)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так уже в библиотеке.\n", b.Title)
		return
	}
	b.IsIssued = false
	b.ReaderID = 0
	fmt.Printf("Книга %s возвращена в библиотеку.\n", b.Title)

}
func (b *Book) AssignBooks() {
	fmt.Printf("Читатель с ID %d взял эту книгу: %s by %s\n", b.ReaderID, b.Title, b.Author)
}

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

}
