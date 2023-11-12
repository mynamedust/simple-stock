// Приложение для взаимодействия с товарами на складах.
package main

import (
	"errors"
	"github.com/mynamedust/simple-stock/internal/config"
	"github.com/mynamedust/simple-stock/internal/server"
	"log"
	"net/http"
)

func main() {
	сfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	s, err := server.New(сfg)
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
