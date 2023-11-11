package models

type Product struct {
	ID           string `jsonapi:"primary,product"`
	StorehouseID int    `jsonapi:"attr,storehouse_id"`
	Code         string `jsonapi:"attr,code"`
}
