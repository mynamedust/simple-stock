package main

import (
	"errors"
	"lamodaTest/internal/config"
	"lamodaTest/internal/server"
	"log"
	"net/http"
)

func main() {
	if err := server.New(config.New()).Run(); !errors.Is(err, http.ErrServerClosed) {
		log.Println(err)
		return
	}
}
