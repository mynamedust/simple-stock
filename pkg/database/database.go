// Package database Реализует работу приложения с базой данных
// postgres, а так же определяет типы и интерфейсы для взаимодействия.
package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mynamedust/simple-stock/pkg/models"
)

type dataStorage struct {
	database *sql.DB
}

// New Конструктор конфигурации сервера.
func New(cfg models.StorageConfig) (*dataStorage, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSL))
	if err != nil {
		return nil, err
	}
	return &dataStorage{
		database: db,
	}, nil
}

// Close Закрытие соединения с базой данных.
func (db *dataStorage) Close() {
	db.Close()
}

// GetStorehouseRemainderByID Вернуть количество товаров на складе по ID.
func (db *dataStorage) GetStorehouseRemainderByID(id int) (int, error) {
	var count int

	err := db.database.QueryRow("SELECT COALESCE(SUM(quantity), 0) FROM product WHERE stock_id = $1", id).Scan(&count)
	return count, err
}

// ReserveProductsByCode Резервация товаров по уникальному коду и ID склада.
func (db *dataStorage) ReserveProductsByCode(transaction StockTransaction) error {
	//*ПОДУМАТЬ НАД ЛОГИРОВАНИЕМ*
	//ПРОВЕРЯЕМ СКЛАДЫ
	var count int
	tx, err := db.database.Begin()
	if err != nil {
		return err
	}

	query := "SELECT COUNT(*) FROM storehouse WHERE is_available = true AND id IN ("
	for _, stock := range transaction.Stocks {
		query += fmt.Sprintf("%d,", stock)
	}
	query = query[:len(query)-1] + ")"
	err = tx.QueryRow(query).Scan(&count)
	if err != nil {
		tx.Rollback()
		return err
	}
	if count != len(transaction.Stocks) {
		return errors.New("Storehouses are not available")
	}

	//ВЫЧИТАЕМ ИЗ ОБЩЕГО КОЛИЧЕСТВА
	query = fmt.Sprintf(`
        UPDATE product 
        SET quantity = quantity - (
            SELECT COUNT(*)
            FROM (VALUES %s) AS v(code, stock_id)
            WHERE v.code = product.code AND v.stock_id = product.stock_id
        )
        WHERE (code, stock_id) IN (%s)
    `, transaction.Products, transaction.Products)

	result, err := tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if int(rows) != transaction.Count {
		tx.Rollback()
		return errors.New("Not all products are in storehouse")
	}

	//ДОБАВЛЯЕМ РЕЗЕРВАЦИЮ
	query = fmt.Sprintf(`
        UPDATE product 
        SET reserved = reserved + (
            SELECT COUNT(*)
            FROM (VALUES %s) AS v(code, stock_id)
            WHERE v.code = product.code AND v.stock_id = product.stock_id
        )
        WHERE (code, stock_id) IN (%s)
    `, transaction.Products, transaction.Products)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// ReleaseProductsByCode Освобождение резерва товаров по уникальному коду и ID склада.
func (db *dataStorage) ReleaseProductsByCode(transaction StockTransaction) error {
	var count int
	tx, err := db.database.Begin()
	if err != nil {
		return err
	}

	query := "SELECT COUNT(*) FROM storehouse WHERE is_available = true AND id IN ("
	for _, stock := range transaction.Stocks {
		query += fmt.Sprintf("%d,", stock)
	}
	query = query[:len(query)-1] + ")"
	err = tx.QueryRow(query).Scan(&count)
	if err != nil {
		tx.Rollback()
		return err
	}
	if count != len(transaction.Stocks) {
		return errors.New("Storehouses are not available")
	}

	//ВЫЧИТАЕМ ИЗ РЕЗЕРВАЦИИ
	query = fmt.Sprintf(`
        UPDATE product 
        SET reserved = reserved - (
            SELECT COUNT(*)
            FROM (VALUES %s) AS v(code, stock_id)
            WHERE v.code = product.code AND v.stock_id = product.stock_id
        )
        WHERE (code, stock_id) IN (%s)
    `, transaction.Products, transaction.Products)

	result, err := tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if int(rows) != transaction.Count {
		tx.Rollback()
		return errors.New("Not all products are in storehouse")
	}

	//ДОБАВЛЯЕМ В ОБЩЕЕ КОЛИЧЕСТВО
	query = fmt.Sprintf(`
        UPDATE product 
        SET quantity = quantity + (
            SELECT COUNT(*)
            FROM (VALUES %s) AS v(code, stock_id)
            WHERE v.code = product.code AND v.stock_id = product.stock_id
        )
        WHERE (code, stock_id) IN (%s)
    `, transaction.Products, transaction.Products)

	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
