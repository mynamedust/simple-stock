package interfaces

import (
	"context"
	"github.com/mynamedust/simple-stock/pkg/models"
)

// Storage Интерфейс, реализующий резервацию и освобождение резерва товара.
type Storage interface {
	Reserver
	Releaser
	StorehouseRemainderGetter
	Close()
}

type Reserver interface {
	ReserveProducts(ctx context.Context, quantityByID map[int]int, storehouseID int) error
}

type Releaser interface {
	ReleaseProducts(ctx context.Context, quantityByID map[int]int, storehouseID int) error
}

type StorehouseRemainderGetter interface {
	GetStorehousesRemainder(ctx context.Context, storehouseID int) ([]models.Product, error)
}
