// Package server Пакет предоставляет функции для запуска и настройки веб-сервера.
// Он включает в себя конструктор, роутер(github.com/gorilla/mux) и функции-обработчики,
// а также зависит от конфигурации, предоставляемой пакетом "simple-stock/internal/config".
package server

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"simple-stock/internal/config"
	"simple-stock/pkg/database"
)

// Server Структура веб-сервера.
type Server struct {
	router  *mux.Router
	storage database.Storage
	address string //`env:"SERVER_ADDRESS"`
	logger  *zap.SugaredLogger
}

// Run Функция запуска веб-сервера.
func (s *Server) Run() error {
	defer s.logger.Sync()
	defer s.storage.Close()

	s.logger.Infow(
		"Starting server",
		"Address", s.address)
	return http.ListenAndServe(s.address, s.router)
}

// New Конструктор сервера управления товарами.
func New(cfg config.Server) (s *Server) {
	var err error
	s = &Server{
		address: cfg.Address,
		storage: database.New(cfg.Database),
	}

	s.initRouter()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.initLogger(); err != nil {
		log.Fatal(err)
	}
	return s
}
