package storage

type Storable interface {
	Save(filename string) error
	Load(filename string) error
}

/*Library является лучшей, она отвечат за хранение и загрузку новых пользователей или читателей.*/
