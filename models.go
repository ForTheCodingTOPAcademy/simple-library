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

//Методы книг

func (b *Book) String() string {
	id := 0
	if b.IsIssued {
		id = b.ReaderID
	}
	return fmt.Sprintf("-----------------------------------\n%d: %s %s by %s. Выдана: %v. ID пользователя: %d\n-----------------------------------", b.ID, b.Title, b.Author.Name, b.Author.Surname, b.IsIssued, id)
}

func (b *Book) IssueBook(r Reader) error {
	if b.IsIssued {
		return fmt.Errorf("Книга %s by %s %s уже выдана другому читателю \n", b.Title, b.Author.Name, b.Author.Surname)
	}
	if !r.IsActive {
		return fmt.Errorf("Читатель %s неактивен и не может получить книгу\n", r.Name)
	}
	b.ReaderID = r.ID
	b.IsIssued = true
	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {

		return fmt.Errorf("Книга %s by %s %s и так уже в библиотеке.\n", b.Title, b.Author.Name, b.Author.Surname)
	}
	b.IsIssued = false
	b.ReaderID = 0
	fmt.Printf("Книга %s by %s %s возвращена в библиотеку.\n\n", b.Title, b.Author.Name, b.Author.Surname)
	return nil
}

//Методы читателей

func (r *Reader) String() string {
	return fmt.Sprintf("-----------------------------------\n%d: %s. Активен: %v\n-----------------------------------", r.ID, r.Name, r.IsActive)
}

//Методы самой библиотеки

//1. Методы вывода

func (l *Library) GetAllBooks() {
	fmt.Print("\n\n==========Все книги в начилии: ==========\n")
	for _, value := range l.Books {
		fmt.Println(value)
	}
	fmt.Print("=========================================\n\n")
}

func (l *Library) GetAllReaders() {
	fmt.Print("\n\n==========Все зарегестрированные чители: ==========\n")
	for _, value := range l.Readers {
		fmt.Println(value)
	}
	fmt.Print("==================================================\n\n")
}

//2.Методы поиска

// Поиск книги по ID
func (l *Library) GetBookById(BookId int) (*Book, error) {
	for index := range l.Books {
		if BookId == l.Books[index].ID {
			fmt.Println("Книга найдена: \n", l.Books[index].String())
			return &l.Books[index], nil
		}
	}
	return nil, fmt.Errorf("\n\nКнига не найдена! Просмотрите список всех книг и найдите нужный Вам ID: GetAllBooks()\n\n")
}

// Поиск книги по названию
func (l *Library) GetBookByTitle(BookTitle string) (*Book, error) {
	for index := range l.Books {
		if l.Books[index].Title == BookTitle {
			fmt.Println("Книга найдена: \n", l.Books[index].String())
			return &l.Books[index], nil
		}
	}
	return nil, fmt.Errorf("\n\nКнига не найдена! Просмотрите список всех книг и найдите нужное Вам название: GetAllBooks()\n\n")
}

// Поиск читателя по ID
func (l *Library) GetReaderById(ReaderID int) (*Reader, error) {
	for index := range l.Readers {
		if l.Readers[index].ID == ReaderID {
			fmt.Println("Читатель найден: \n", l.Readers[index].String())
			return &l.Readers[index], nil
		}
	}
	return nil, fmt.Errorf("\n\nЧитатель не найден! Просмотрите список всех читателей и найдите нужный Вам ID: GetAllReaders()\n\n")
}

// Поиск читателя по имени
func (l *Library) GetReaderByName(ReaderName string) (*Reader, error) {
	for index := range l.Readers {
		if l.Readers[index].Name == ReaderName {
			fmt.Println("Читатель найден: \n", l.Readers[index].String())
			return &l.Readers[index], nil
		}
	}
	return nil, fmt.Errorf("\n\nЧитатель не найден! Просмотрите список всех читателей и найдите нужное Вам Имя: GetAllReaders()\n\n")
}

//3.Методы добавления

func (l *Library) AddBook(title string, author_name string, author_surname string) (*Book, error) {
	if _, e := l.GetBookByTitle(title); e != nil {
		return nil, fmt.Errorf("Книга с таким названием уже есть в библиотеке.")
	}
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
	return &book, nil

}

// 4.Связующие методы

func (l *Library) IssueBookToReader(BookID int, ReaderID int) error {
	var book *Book
	var reader *Reader
	var err error
	fmt.Print("\n\n======Выдача книги: ======\n\n")
	if book, err = l.GetBookById(BookID); err != nil {
		return fmt.Errorf("Некорректный ID книги!")
	}
	if reader, err = l.GetReaderById(ReaderID); err != nil {
		return fmt.Errorf("Некорректный ID читателя!")
	}
	if eror := book.IssueBook(*reader); eror != nil {
		return eror
	}

	fmt.Printf("Книга %s by %s %s выдана %s\n", book.Title, book.Author.Name, book.Author.Surname, reader.Name)
	fmt.Print("\n==========================\n\n")
	return nil
}

func (l *Library) ReturnBook(BookID int) error {
	b, err := l.GetBookById(BookID)

	if err != nil {
		return err
	}

	err = b.ReturnBook()

	if err != nil {
		return err
	}
	return nil
}
