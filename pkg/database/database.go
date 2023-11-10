package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"lamodaTest/internal/config"
	"log"
)

type Storage interface {
	GetStockRemainderByID(id int) (count int, err error)
	ReserveProductsByCode(products string) (err error)
	ReleaseProductsByCode(products string) (err error)
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

func (db *DataStorage) ReserveProductsByCode(products string) (err error) {
	query := `UPDATE product AS p
		SET reserved = CASE
			WHEN s.is_available = true AND p.reserved + 1 <= p.quantity THEN p.reserved + 1
			ELSE p.reserved
		END
		FROM storehouse AS s
		WHERE p.stock_id = s.id AND (p.code, p.stock_id) IN (VALUES ` + products + `)`
	_, err = db.database.Exec(query)
	return
}

func (db *DataStorage) ReleaseProductsByCode(products string) (err error) {
	query := `UPDATE product AS p
		SET reserved = CASE
			WHEN s.is_available = true AND p.reserved - 1 >= 0 THEN p.reserved - 1
			ELSE p.reserved
		END
		FROM storehouse AS s
		WHERE p.stock_id = s.id AND (p.code, p.stock_id) IN (VALUES ` + products + `)`
	_, err = db.database.Exec(query)
	return
}
