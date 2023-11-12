package database

// Storage Интерфейс, реализующий резервацию и освобождение резерва товара.
type Storage interface {
	Reserver
	Releaser
	Counter
	Close()
}

// Reserver Интерфейс, реализующий резервацию товара.
type Reserver interface {
	ReserveProductsByCode(transaction StockTransaction) error
}

// Releaser Интерфейс, реализующий освобождение резерва товара.
type Releaser interface {
	ReleaseProductsByCode(transaction StockTransaction) error
}

// Counter Интерфейс возврата количества товара на складе.
type Counter interface {
	GetStorehouseRemainderByID(id int) (count int, err error)
}

// StockTransaction Структура данных резервации/освобождения резервации.
type StockTransaction struct {
	Products string
	Stocks   []int
	Count    int
}
