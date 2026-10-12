package storage

import (
	"os"
	"encoding/json"
	"library-project/domain"
	"library-project/library"
)

type Storable interface {
	Save(filename string) error
	Load(filename string) error
}

//Вспомагательная структура для хранения
type storageData struct{
	books []*domain.Book `json:"books"`//Обозначение ключа для json
	readers []*domain.Reader `json:"readers"`//Обозначение ключа для json
}

//Метод сохранения библиотеки в JSON
func SaveLibraryToJSON(filePath string, lib *library.Library) error {
	Data := storageData{books: lib.Books, readers: lib.Readers,}
	
	DataJSON, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
	}
	return os.WriteFile(filePath, DataJSON, 0644)
}

//Метод загрузки библиотеки из JSON
func LoadLibraryFromJSON(filePath string) (*library.Library, error){
	
	DataJSON, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var data storageData

	if err := json.Unmarshal(DataJSON, &data); err != nil {
		return nil, err
	}

	lib := library.CreateLibrary(data.books, data.readers)
	lib.UpdateLastIDs()

	return lib, nil
}


