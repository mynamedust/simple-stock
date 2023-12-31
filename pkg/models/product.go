// Package models описывает модели.
package models

// Product модель товара.
type Product struct {
	ID           int    `jsonapi:"primary,product",gorm:"primaryKey"`
	StorehouseID int    `jsonapi:"attr,storehouse_id",gorm:"column:storehouse_id"`
	Size         string `jsonapi:"attr,size",gorm:"column:size"`
	Quantity     int    `jsonapi:"attr,quantity",gorm:"column:quantity"`
	Reserved     int    `jsonapi:"attr,reserved",gorm:"column:reserved"`
	Name         string `jsonapi:"attr,name",gorm:"column:name"`
}

// TableName возвращает имя таблицы.
func (Product) TableName() string {
	return "product"
}
