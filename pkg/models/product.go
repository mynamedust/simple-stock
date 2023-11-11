// Package models Пакет предоставляющий модели для товаров и складов.
package models

// Product Модель товаров включающая ID, идентификатор склада и уникальный код.
type Product struct {
	ID           string `jsonapi:"primary,product"`
	StorehouseID int    `jsonapi:"attr,storehouse_id"`
	Code         string `jsonapi:"attr,code"`
}
