package interfaces

import (
	"context"
	"github.com/mynamedust/simple-stock/pkg/models"
)

// Storage содержит сигнатуры резервирования, освобождения резервирования и получения количества товаров на складе.
type Storage interface {
	Reserver
	Releaser
	StorehouseRemainderGetter
	Close()
}

// Reserver содержит сигнатуру метода резервирования товара.
type Reserver interface {
	ReserveProducts(ctx context.Context, quantityByID map[int]int, storehouseID int) error
}

// Releaser содержит сигнатуру метода освобождения резервирования товара.
type Releaser interface {
	ReleaseProducts(ctx context.Context, quantityByID map[int]int, storehouseID int) error
}

// StorehouseRemainderGetter содержит сигнатуру метода получения информации о количестве товаров на складе.
type StorehouseRemainderGetter interface {
	GetStorehousesRemainder(ctx context.Context, storehouseID int) ([]models.Product, error)
}
