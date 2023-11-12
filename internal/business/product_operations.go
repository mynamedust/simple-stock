// Package business предоставляет функции для обработки бизнес-логики, связанной с управлением товарами.
// Включает операции по преобразованию товаров в строки и извлечению уникальных идентификаторов складов.
// Функции в этом пакете работают с моделями, определенными в пакете "simple-stock/pkg/models".
package business

import (
	"fmt"
	"github.com/mynamedust/simple-stock/pkg/models"
)

// productToString Функция преобразования слайса товаров в строку и возврата количества уникальных товаров.
func productToString(products []models.Product) (string, int) {
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

// getStocksID Функция извлечения уникальных ID складов.
func getStocksID(products []models.Product) []int {
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
