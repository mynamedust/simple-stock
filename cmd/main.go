// Приложение для взаимодействия с товарами на складах.
package main

import (
	"errors"
	"github.com/mynamedust/simple-stock/internal/config"
	"github.com/mynamedust/simple-stock/internal/server"
	"github.com/mynamedust/simple-stock/pkg/database"
	"log"
	"net/http"
)

func main() {
	// Создание конфигурации приложения
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	// Инициализация базы данных
	storage, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	s, err := server.New(cfg, storage)
	if err != nil {
		log.Fatal(err)
	}

	// Запуск приложения
	if err = s.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
