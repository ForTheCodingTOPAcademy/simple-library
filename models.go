package main

import (
	"fmt"
)

type Book struct {
	ID       int
	Title    string
	Author   Author
	IsIssued bool
	ReaderID int
}

type Reader struct {
	ID       int
	Name     string
	IsActive bool
}

type Author struct {
	Name    string
	Surname string
}

type Library struct {
	Readers []Reader
	Books   []Book

	lastAddedBookID   int
	lastAddedReaderID int
}

func (b *Book) String() string {
	id := 0
	if b.IsIssued {
		id = b.ReaderID
	}
	return fmt.Sprintf("-----------------------------------\n%d: %s %s by %s. Выдана: %v. ID пользователя: %d\n-----------------------------------", b.ID, b.Title, b.Author.Name, b.Author.Surname, b.IsIssued, id)
}

func (r *Reader) String() string {
	return fmt.Sprintf("-----------------------------------\n%d: %s. Активен: %v\n-----------------------------------", r.ID, r.Name, r.IsActive)
}

/*В разработке
func (b *Book) IssueBook(r Reader) {
	if b.IsIssued {
		fmt.Printf("Книга %s by %s %s уже кому-то выдана читателю \n", b.Title, b.Author.Name, b.Author.Surname)
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

}*/

func (l *Library) AddBook(title string, author_name string, author_surname string) *Book {
	book := Book{
		ID:       l.lastAddedBookID + 1,
		Title:    title,
		Author:   Author{Name: author_name, Surname: author_surname},
		IsIssued: false,
		ReaderID: 0,
	}
	l.Books = append(l.Books, book)
	l.lastAddedBookID++
	fmt.Printf("\n\nДобавленная книга: \n%v\n\n", book.String())
	return &book
}

func (l *Library) GetAllBooks() {
	for _, value := range l.Books {
		fmt.Println(value.String())
	}
}

// Поиск книги по ID
func (l *Library) GetBookById(BookId int) *Book {
	for _, book := range l.Books {
		if BookId == book.ID {
			fmt.Println("Книга найдена: \n", book.String())
			return &book
		}
	}
	fmt.Printf("\n\nКнига не найдена! Просмотрите список всех книг и найдите нужный Вам ID: GetAllBooks()\n\n")
	return nil
}

// Поиск книги по названию
func (l *Library) GetBookByTitle(BookTitle string) *Book {
	for _, book := range l.Books {
		if book.Title == BookTitle {
			fmt.Println("Книга найдена: \n", book.String())
			return &book
		}
	}
	fmt.Printf("\n\nКнига не найдена! Просмотрите список всех книг и найдите нужное Вам название: GetAllBooks()\n\n")
	return nil
}

// Поиск читателя по ID
func (l *Library) GetReaderById(ReaderID int) *Reader {
	for _, reader := range l.Readers {
		if reader.ID == ReaderID {
			fmt.Println("Читатель найден: \n", reader.String())
			return &reader
		}
	}
	fmt.Printf("\n\nЧитатель не найден! Просмотрите список всех читателей и найдите нужный Вам ID: GetAllReaders()\n\n")
	return nil
}

// Поиск читателя по имени
func (l *Library) GetReaderByName(ReaderName string) *Reader {
	for _, reader := range l.Readers {
		if reader.Name == ReaderName {
			fmt.Println("Читатель найден: \n", reader.String())
			return &reader
		}
	}
	fmt.Printf("\n\nЧитатель не найден! Просмотрите список всех читателей и найдите нужное Вам Имя: GetAllReaders()\n\n")
	return nil
}
