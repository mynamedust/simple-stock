package models

type Product struct {
	ID       int
	Name     string
	Size     string
	Code     string
	Quantity int
	Reserved int
	StockID  int
}
