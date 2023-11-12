package business

import (
	"github.com/mynamedust/simple-stock/pkg/database"
	"github.com/mynamedust/simple-stock/pkg/models"
)

func createTransaction(products []models.Product) database.StockTransaction {
	productsString, count := productToString(products)
	stocksID := getStocksID(products)

	return database.StockTransaction{
		Products: productsString,
		Count:    count,
		Stocks:   stocksID,
	}
}

func Reserve(products []models.Product, storage database.Reserver) error {
	transaction := createTransaction(products)
	return storage.ReserveProductsByCode(transaction)
}

func Release(products []models.Product, storage database.Releaser) error {
	transaction := createTransaction(products)
	return storage.ReleaseProductsByCode(transaction)
}

func GetStock(id int, storage database.Counter) (int, error) {
	return storage.GetStorehouseRemainderByID(id)
}
