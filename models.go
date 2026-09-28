package main

import (
	"fmt"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	IsIssued bool
}

type Reader struct {
	ID       int
	Name     string
	IsActive bool
}

func (b Book) String() string {
	return fmt.Sprintf("-----------------------------------\n%d: %s by %s. Выдана: %v\n-----------------------------------", b.ID, b.Title, b.Author, b.IsIssued)
}

func (b *Book) IssueBook(r Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s by %s уже кому-то выдана\n", b.Title, b.Author)
		return
	}
	if !r.IsActive {
		fmt.Printf("Читатель %s неактивен и не может получить книгу\n", r.Name)
		return
	}
	b.IsIssued = true
	fmt.Printf("Книга %s выдана %s\n", b.Title, r.Name)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так уже в библиотеке.\n", b.Title)
		return
	}
	fmt.Printf("Книга %s возвращена в библиотеку.\n", b.Title)

}

func main() {
	b := Book{ID: 1, Title: "Вокруг света за 80 дней", Author: "Жуль Верн", IsIssued: false}
	r := Reader{ID: 10, Name: "Саня", IsActive: false}
	fmt.Println(b.String())
	b.ReturnBook()
	b.IssueBook(r)
	r.IsActive = true
	b.IssueBook(r)
	b.ReturnBook()
	fmt.Println(b.String())

}
