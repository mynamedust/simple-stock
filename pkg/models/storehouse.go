package models

// Storehouse Модель складов включающая идентификатор склада.
type Storehouse struct {
	ID int `jsonapi:"primary,storehouse"`
}
