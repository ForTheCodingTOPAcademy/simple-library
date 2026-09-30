package main

import (
	"fmt"
)

func main() {
	//Создание без функций
	readers := []Reader{
		Reader{ID: 1, Name: "Мигель", IsActive: false},
		Reader{ID: 2, Name: "Александр", IsActive: false},
		Reader{ID: 3, Name: "Анатолий", IsActive: false},
		Reader{ID: 4, Name: "Милана", IsActive: false},
		Reader{ID: 5, Name: "Игорь", IsActive: false},
		Reader{ID: 6, Name: "Алана", IsActive: false},
		Reader{ID: 7, Name: "Виктория", IsActive: false},
	}
	books := []Book{
		Book{ID: 1, Title: "Вокруг света за 80 дней", Author: Author{Name: "Жуль", Surname: "Верн"}, IsIssued: false},
		Book{ID: 2, Title: "Таинственный остров", Author: Author{Name: "Жуль", Surname: "Верн"}, IsIssued: false},
		Book{ID: 3, Title: "Остров сокровищ", Author: Author{Name: "Жуль", Surname: "Верн"}, IsIssued: false},
		Book{ID: 4, Title: "2:36 по Аляске", Author: Author{Name: "Анастасия", Surname: "Гор"}, IsIssued: false},
		Book{ID: 5, Title: "Менталитет Мамбы", Author: Author{Name: "Коби", Surname: "Брайнт"}, IsIssued: false},
	}

	library := Library{Readers: readers, Books: books}
	library.lastAddedBookID = 5
	library.lastAddedReaderID = 7

	//Работа с функциями
	library.GetAllBooks()
	library.AddBook("Капитанская дочка", "Александр", "Пушкин")
	library.GetAllBooks()

	existing_book_by_ID, err := library.GetBookById(1)
	non_existring_book_by_ID, err := library.GetBookById(10)
	existing_book_by_Title, err := library.GetBookByTitle("2:36 по Аляске")
	non_existring_book_by_Title, err := library.GetBookByTitle("Война и мир")

	existing_reader_by_ID, err := library.GetReaderById(1)
	non_existring_reader_by_ID, err := library.GetReaderById(10)
	existing_reader_by_Name, err := library.GetReaderByName("Виктория")
	non_existring_reader_by_Name, err := library.GetReaderByName("Сергей")

	existing_reader_by_ID.IsActive = true

	fmt.Println("\n\n\nВыдача книги\n\n\n")

	if eror := library.IssueBookToReader(1, 1); eror != nil {

	} else {
		fmt.Println(eror)
	}

	fmt.Println(existing_reader_by_ID.IsActive)

	//Пока возвращаемые ID не используется, поэтому я задействовал их тут, чтобы программа не выдавала ошибку о неиспользованных переменных
	fmt.Sprintln(err, existing_book_by_ID, existing_book_by_Title, existing_reader_by_ID, existing_reader_by_Name, non_existring_book_by_ID, non_existring_book_by_Title, non_existring_reader_by_ID, non_existring_reader_by_Name)

	/*!!!Все методы расположены в models.go!!!*/

	var emailN Notifier = EmailNotifier{EmailAdress: "example@gmail.com"}
	var SMSN Notifier = SMSNotifier{PhoneNumber: "+7 988 376 88 30"}

	SMSN.Notify("Hello through SMS")
	emailN.Notify("Hello through email")
}
