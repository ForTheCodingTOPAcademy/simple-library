package main

type Storable interface {
	Save() error
	Load() error
}

/*Library является лучшей, она отвечат за хранение и загрузку новых пользователей или читателей.*/
