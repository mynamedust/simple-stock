package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"lamodaTest/internal/config"
	"log"
	"net/http"
)

type Storage interface {
	GetStockRemainderByID(id int) (count int, err error)
	ReserveProductsByCode(stocks []int, products string, count int) (int, error)
	ReleaseProductsByCode(stocks []int, products string, count int) (int, error)
	Close()
}

type DataStorage struct {
	database *sql.DB
}

func New(cfg config.Storage) *DataStorage {
	db, err := sql.Open("postgres", fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSL))
	if err != nil {
		log.Fatal(err)
	}
	return &DataStorage{
		database: db,
	}
}

func (db *DataStorage) Close() {
	db.Close()
}

func (db *DataStorage) GetStockRemainderByID(id int) (count int, err error) {
	err = db.database.QueryRow("SELECT COALESCE(SUM(quantity) - SUM(reserved), 0) FROM product WHERE stock_id = $1", id).Scan(&count)
	return
}

func (db *DataStorage) ReserveProductsByCode(stocks []int, products string, productsCount int) (int, error) {
	//*ПОДУМАТЬ НАД ЛОГИРОВАНИЕМ*
	//ПРОВЕРЯЕМ СКЛАДЫ
	query := "SELECT COUNT(*) FROM storehouse WHERE is_available = true AND id IN ("
	for _, stock := range stocks {
		query += fmt.Sprintf("%d,", stock)
	}
	query = query[:len(query)-1] + ")"

	var count int
	tx, err := db.database.Begin()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = tx.QueryRow(query).Scan(&count)
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}
	if count != len(stocks) {
		return http.StatusOK, errors.New("Storehouses are not available")
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
    `, products, products)

	result, err := tx.Exec(query)
	if err != nil {
		tx.Rollback()
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23514" {
			return http.StatusOK, err
		}
		return http.StatusInternalServerError, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}
	if int(rows) != productsCount {
		tx.Rollback()
		return http.StatusOK, errors.New("Not all products are in storehouse")
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
    `, products, products)
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return http.StatusOK, err
	}
	err = tx.Commit()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (db *DataStorage) ReleaseProductsByCode(stocks []int, products string, productsCount int) (int, error) {
	return 200, nil
}
