package models

type ProductDto struct {
	ID       int `jsonapi:"primary,productdto"`
	Quantity int `jsonapi:"attr,quantity"`
}

type ProductDtoWithStorehouseID struct {
	DtoProductsChangeReservedStatus []*ProductDto `jsonapi:"relation,products"`
	StorehouseID                    int           `jsonapi:"attr,storehouse_id"`
}
