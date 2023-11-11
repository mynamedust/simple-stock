// Приложение для взаимодействия с товарами на складах.
package main

import (
	"errors"
	"log"
	"net/http"
	"simple-stock/internal/config"
	"simple-stock/internal/server"
)

func main() {
	if err := server.New(config.New()).Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
		return
	}
}
