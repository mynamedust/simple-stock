package business

import (
	"fmt"
	"lamodaTest/pkg/models"
)

func ProductToString(products []models.Product) (string, int) {
	var productsSting string
	unique := make(map[models.Product]struct{})

	count := 0
	for _, product := range products {
		if _, exist := unique[product]; !exist {
			count++
		}
		unique[product] = struct{}{}
		productsSting += fmt.Sprintf("('%s', %d),", product.Code, product.StorehouseID)
	}
	// Удаление последней запятой
	return productsSting[:len(productsSting)-1], count
}

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
