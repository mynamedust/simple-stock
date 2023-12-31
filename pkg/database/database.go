// Package database реализует работу приложения с базой данных postgres.
package database

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mynamedust/simple-stock/pkg/models"
	"golang.org/x/exp/maps"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DataStorage представляет собой структуру для взаимодействия с базой данных с использованием GORM.
type DataStorage struct {
	database *gorm.DB
}

// New возвращает указатель на экземпляр DataStorage и ошибку, если подключение неуспешно.
func New(cfg models.StorageConfig) (*DataStorage, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSL)), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DataStorage{
		database: db,
	}, nil
}

// Close закрытие соединения с базой данных.
func (db *DataStorage) Close() {
	db.Close()
}

// GetStorehousesRemainder возвращает информацию о всех товарах лежащих на складе с указанным storehouseID.
func (db *DataStorage) GetStorehousesRemainder(ctx context.Context, storehouseID int) ([]models.Product, error) {
	product := make([]models.Product, 0)

	tx := db.database.WithContext(ctx).Begin()
	if err := tx.Exec(`set transaction isolation level read committed`).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Model(new(models.Product)).
		Select("product.*").
		Joins("INNER JOIN storehouse AS s ON s.id = product.storehouse_id").
		Where("s.id = ?", storehouseID).
		Find(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

// ReserveProducts резервирует товары по ID на складе с указанным storehouseID.
func (db *DataStorage) ReserveProducts(ctx context.Context, quantityByID map[int]int, storehouseID int) error {
	tx := db.database.WithContext(ctx).Begin()

	if err := tx.Exec(`set transaction isolation level serializable`).Error; err != nil {
		tx.Rollback()
		return err
	}
	products := make([]models.Product, 0, len(quantityByID))
	if err := tx.Model(new(models.Product)).
		Select("product.*").
		Joins("INNER JOIN storehouse AS s ON s.id = product.storehouse_id").
		Where("s.is_available = true AND product.id IN ? AND s.id = ?", maps.Keys(quantityByID), storehouseID).
		Find(&products).
		Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(products) != len(quantityByID) {
		tx.Rollback()
		return errors.New("some storehouses are not available")
	}

	for i := range products {
		products[i].Quantity -= quantityByID[products[i].ID]
		if products[i].Quantity < 0 {
			tx.Rollback()
			return errors.New("not enough products in storehouse")
		}
		products[i].Reserved += quantityByID[products[i].ID]
	}
	if err := tx.Save(products).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// ReleaseProducts освобождает зарезервированные товары по ID на складе с указанным storehouseID.
func (db *DataStorage) ReleaseProducts(ctx context.Context, quantityByID map[int]int, storehouseID int) error {
	tx := db.database.WithContext(ctx).Begin()

	if err := tx.Exec(`set transaction isolation level serializable`).Error; err != nil {
		tx.Rollback()
		return err
	}
	products := make([]models.Product, 0, len(quantityByID))
	if err := tx.Model(new(models.Product)).
		Select("product.*").
		Joins("INNER JOIN storehouse AS s ON s.id = product.storehouse_id").
		Where("s.is_available = true AND product.id IN ? AND s.id = ?", maps.Keys(quantityByID), storehouseID).
		Find(&products).
		Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(products) != len(quantityByID) {
		tx.Rollback()
		return errors.New("some storehouses are not available")
	}

	for i := range products {
		products[i].Reserved -= quantityByID[products[i].ID]
		if products[i].Reserved < 0 {
			tx.Rollback()
			return errors.New("not enough reserved products in storehouse")
		}
		products[i].Quantity += quantityByID[products[i].ID]
	}

	if err := tx.Save(products).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
