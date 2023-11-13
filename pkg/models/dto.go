package models

// ProductDto представляет структуру данных для передачи информации о продукте в формате JSON:API.
type ProductDto struct {
	ID       int `jsonapi:"primary,product_dto"`
	Quantity int `jsonapi:"attr,quantity"`
}

// StorehouseDto представляет структуру данных для передачи информации о складе в формате JSON:API.
type StorehouseDto struct {
	ID          int           `jsonapi:"primary,storehouse_dto"`
	ProductsDto []*ProductDto `jsonapi:"relation,products"`
}
