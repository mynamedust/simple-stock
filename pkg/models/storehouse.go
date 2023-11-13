package models

// Storehouse модель склада.
type Storehouse struct {
	ID    int `jsonapi:"primary,storehouse"`
	Count int `jsonapi:"attr,count,omitempty"`
}

// TableName возвращает имя таблицы.
func (Storehouse) TableName() string {
	return "storehouse"
}
