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

func Reserve(products models.ProductDtoWithStorehouseID, storage interfaces.Reserver, log *zap.SugaredLogger) error {
	quantityByID := make(map[int]int, len(products.DtoProductsChangeReservedStatus))

	for _, product := range products.DtoProductsChangeReservedStatus {
		quantityByID[product.ID] = product.Quantity
	}
	log.Debugw("reserve products",
		"quantityByID", quantityByID,
		"storehouseID", products.StorehouseID)

	if err := storage.ReserveProducts(context.Background(), quantityByID, products.StorehouseID); err != nil {
		return errors.Wrap(err, "failed to reserve products")
	}

	return nil
}

func Release(products models.ProductDtoWithStorehouseID, storage interfaces.Releaser, log *zap.SugaredLogger) error {
	quantityByID := make(map[int]int, len(products.DtoProductsChangeReservedStatus))

	for _, product := range products.DtoProductsChangeReservedStatus {
		quantityByID[product.ID] = product.Quantity
	}
	log.Debugw("reserve products",
		"quantityByID", quantityByID,
		"storehouseID", products.StorehouseID)

	if err := storage.ReleaseProducts(context.Background(), quantityByID, products.StorehouseID); err != nil {
		return errors.Wrap(err, "failed to release products")
	}

	return nil
}

func GetRemainder(id int, storage interfaces.StorehouseRemainderGetter, log *zap.SugaredLogger) (int, error) {
	log.Debugw("get storehouse remainder")

	products, err := storage.GetStorehousesRemainder(context.Background(), id)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get storehouse remainder")
	}

	count := 0
	for _, product := range products {
		count += product.Quantity
	}
	return count, nil
}
