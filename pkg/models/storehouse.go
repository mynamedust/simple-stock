package models

// Storehouse Модель складов включающая идентификатор склада.
type Storehouse struct {
	ID    int `jsonapi:"primary,storehouse"`
	Count int `jsonapi:"attr,count,omitempty"`
}

func (Storehouse) TableName() string {
	return "storehouse"
}
