// Package business предоставляет функции для обработки бизнес-логики, связанной с управлением товарами.
// Включает операции по резервации, освобождению резерва и получению количества товара на складе.
// Функции в этом пакете работают с моделями, определенными в пакете "simple-stock/pkg/models".
package business

import (
	"context"
	"github.com/mynamedust/simple-stock/internal/interfaces"
	"github.com/mynamedust/simple-stock/pkg/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Reserve Функция резервации товаров на складе.
func Reserve(products models.StorehouseDto, storage interfaces.Reserver, log *zap.SugaredLogger) error {
	quantityByID := make(map[int]int, len(products.ProductsDto))

	for _, product := range products.ProductsDto {
		quantityByID[product.ID] = product.Quantity
	}
	log.Debugw("reserve products",
		"quantityByID", quantityByID,
		"storehouseID", products.ID)

	if err := storage.ReserveProducts(context.Background(), quantityByID, products.ID); err != nil {
		return errors.Wrap(err, "failed to reserve products")
	}

	return nil
}

// Release Функция освобождения резервации на складе по уникальному коду товаров.
func Release(products models.StorehouseDto, storage interfaces.Releaser, log *zap.SugaredLogger) error {
	quantityByID := make(map[int]int, len(products.ProductsDto))

	for _, product := range products.ProductsDto {
		quantityByID[product.ID] = product.Quantity
	}
	log.Debugw("reserve products",
		"quantityByID", quantityByID,
		"storehouseID", products.ID)

	if err := storage.ReleaseProducts(context.Background(), quantityByID, products.ID); err != nil {
		return errors.Wrap(err, "failed to release products")
	}

	return nil
}

// GetStorehouseRemainder Получение информации о товарах на складе.
func GetStorehouseRemainder(id int, storage interfaces.StorehouseRemainderGetter, log *zap.SugaredLogger) (*[]*models.ProductDto, error) {
	log.Debugw("get storehouse remainder")

	products, err := storage.GetStorehousesRemainder(context.Background(), id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get storehouse remainder")
	}

	dto := make([]*models.ProductDto, 0, len(products))
	for _, product := range products {
		dto = append(dto, &models.ProductDto{ID: product.StorehouseID, Quantity: product.Quantity})
	}
	return &dto, nil
}
