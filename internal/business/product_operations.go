// Package business предоставляет функции для обработки бизнес-логики, связанной с управлением товарами.
// Включает операции по преобразованию товаров в строки и извлечению уникальных идентификаторов складов.
// Функции в этом пакете работают с моделями, определенными в пакете "simple-stock/pkg/models".
package business

import (
	"fmt"
	"simple-stock/pkg/models"
)

// ProductToString Функция преобразования слайса товаров в строку и возврата количества уникальных товаров.
func ProductToString(products []models.Product) (string, int) {
	var productsSting string

	type uniqueProduct struct {
		storehouseID int
		code         string
	}
	unique := make(map[uniqueProduct]struct{})

	count := 0
	for _, product := range products {
		currProduct := uniqueProduct{storehouseID: product.StorehouseID, code: product.Code}
		if _, exist := unique[currProduct]; !exist {
			count++
		}
		unique[currProduct] = struct{}{}
		productsSting += fmt.Sprintf("('%s', %d),", product.Code, product.StorehouseID)
	}
	// Удаление последней запятой
	return productsSting[:len(productsSting)-1], count
}

// GetStocksID Функция извлечения уникальных ID складов.
func GetStocksID(products []models.Product) []int {
	uniqueStocks := make(map[int]struct{})
	for _, product := range products {
		uniqueStocks[product.StorehouseID] = struct{}{}
	}

	var stocks []int
	for id := range uniqueStocks {
		stocks = append(stocks, id)
	}

	return stocks
}
