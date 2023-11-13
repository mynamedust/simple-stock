// Package server Пакет предоставляет функции для запуска и настройки веб-сервера.
// Он включает в себя конструктор, роутер(github.com/gorilla/mux) и функции-обработчики,
// а также зависит от конфигурации, предоставляемой пакетом "simple-stock/internal/config".
package server

import (
	"github.com/gorilla/mux"
	"github.com/mynamedust/simple-stock/internal/interfaces"
	"github.com/mynamedust/simple-stock/pkg/models"
	"go.uber.org/zap"
	"net/http"
)

// Server Структура веб-сервера.
type Server struct {
	router  *mux.Router
	storage interfaces.Storage
	address string
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

// New возвращает указатель на экземпляр сервера и ошибку, если она есть.
func New(cfg models.ServerConfig, storage interfaces.Storage) (*Server, error) {
	s := &Server{
		address: cfg.Address,
		storage: storage,
	}

	// Инициализация роутинга
	s.router = mux.NewRouter()
	s.router.Use(s.contentTypeCheck)
	s.router.HandleFunc("/products/stock", s.GetStorehouseRemainder).Methods("GET")
	s.router.HandleFunc("/products/reserve", s.reserveProducts).Methods("POST")
	s.router.HandleFunc("/products/release", s.releaseProducts).Methods("POST")

	// Инициализация логгера
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	s.logger = logger.Sugar()

	return s, nil
}
