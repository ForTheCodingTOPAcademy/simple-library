package main

import (
	"flag"
	"fmt"
	"library-project/domain"
	"library-project/library"
	"library-project/storage"
	"log"
	//"library-project/notifications"
)

func main() {
	//Работа с флагами комндной строки

	saveFlag := flag.Bool("save", false, "saving books flag")
	loadFlag := flag.Bool("load", false, "loading books flag")

	flag.Parse()

	//Заполнение библиотеки без функций

	readers := []domain.Reader{
		domain.Reader{ID: 1, Name: "Мигель", IsActive: false},
		domain.Reader{ID: 2, Name: "Александр", IsActive: false},
		domain.Reader{ID: 3, Name: "Анатолий", IsActive: false},
		domain.Reader{ID: 4, Name: "Милана", IsActive: false},
		domain.Reader{ID: 5, Name: "Игорь", IsActive: false},
		domain.Reader{ID: 6, Name: "Алана", IsActive: false},
		domain.Reader{ID: 7, Name: "Виктория", IsActive: false},
	}
	books := []domain.Book{
		domain.Book{ID: 1, Title: "Вокруг света за 80 дней", Author: domain.Author{Name: "Жуль", Surname: "Верн"}, IsIssued: false},
		domain.Book{ID: 2, Title: "Таинственный остров", Author: domain.Author{Name: "Жуль", Surname: "Верн"}, IsIssued: false},
		domain.Book{ID: 3, Title: "Остров сокровищ", Author: domain.Author{Name: "Жуль", Surname: "Верн"}, IsIssued: false},
		domain.Book{ID: 4, Title: "2:36 по Аляске", Author: domain.Author{Name: "Анастасия", Surname: "Гор"}, IsIssued: false},
		domain.Book{ID: 5, Title: "Менталитет Мамбы", Author: domain.Author{Name: "Коби", Surname: "Брайнт"}, IsIssued: false},
	}
	fmt.Sprint(books,readers)
	//library := library.CreateLibrary([]domain.Book{}, []domain.Reader{})

	lib := library.CreateLibrary([]domain.Book{}, []domain.Reader{})
	var errr error

	if *loadFlag{
		lib, errr = storage.LoadLibraryFromJSON("books.json")
		if errr != nil {
			log.Fatal(errr)
		}
	}else{
		lib = library.CreateLibrary(books, readers)
	}



	fmt.Println("Книг в библиотеке: ", lib.GetLastAddedBookID()+1)
	fmt.Println("Читателей в библиотеке: ", lib.GetLastAddedReaderID()+1)

	//Заполнение библиотеки с функциями
	lib.GetAllBooks()
	_, er := lib.AddBook("Капитанская дочка", "Александр", "Пушкин")
	if er != nil {
		fmt.Println(er)
	}
	lib.GetAllBooks()

	//Проверка работы функций по поиску книг и читателей

	existing_book_by_ID, err := lib.GetBookById(1)
	non_existring_book_by_ID, err := lib.GetBookById(10)
	existing_book_by_Title, err := lib.GetBookByTitle("2:36 по Аляске")
	non_existring_book_by_Title, err := lib.GetBookByTitle("Война и мир")

	existing_reader_by_ID, err := lib.GetReaderById(1)
	non_existring_reader_by_ID, err := lib.GetReaderById(10)
	existing_reader_by_Name, err := lib.GetReaderByName("Виктория")
	non_existring_reader_by_Name, err := lib.GetReaderByName("Сергей")

	//Проверка работы функций по выдаче книг читателю

	if eror := lib.IssueBookToReader(1, 1); eror != nil {
		fmt.Println(eror)
	}

	existing_reader_by_ID.IsActive = true

	if eror := lib.IssueBookToReader(1, 1); eror != nil {
		fmt.Println(eror)
	}

	//Проверка работы функций по возврате книг в библиотеку

	e := lib.ReturnBook(existing_book_by_Title.ID)
	if e != nil {
		fmt.Println(e)
	}
	e = lib.ReturnBook(existing_book_by_ID.ID)
	if e != nil {
		fmt.Println(e)
	}

	//Пока возвращаемые ID не используется, поэтому я задействовал их тут, чтобы программа не выдавала ошибку о неиспользованных переменных
	fmt.Sprintln(err, existing_book_by_ID, existing_book_by_Title, existing_reader_by_ID, existing_reader_by_Name, non_existring_book_by_ID, non_existring_book_by_Title, non_existring_reader_by_ID, non_existring_reader_by_Name)

	/*!!!Все методы и классы библиотеки расположены в models.go!!!*/

	/*var emailN NoNotifier = EmailNotifier{EmailAdress: "example@gmail.com"}
	var SMSN Notifier = SMSNotifier{PhoneNumber: "+7 988 376 88 30"}

	SMSN.Notify("Hello through SMS")
	emailN.Notify("Hello through email")*/

	
	if *saveFlag {
		if err := storage.SaveLibraryToJSON("books.json", lib); err != nil {
			log.Fatal(err)
		}
	}
}
