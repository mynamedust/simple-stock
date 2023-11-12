// Package server Пакет предоставляет функции для запуска и настройки веб-сервера.
// Он включает в себя конструктор, роутер(github.com/gorilla/mux) и функции-обработчики,
// а также зависит от конфигурации, предоставляемой пакетом "simple-stock/internal/config".
package server

import (
	"github.com/gorilla/mux"
	"github.com/mynamedust/simple-stock/pkg/database"
	"github.com/mynamedust/simple-stock/pkg/models"
	"go.uber.org/zap"
	"net/http"
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
func New(cfg models.ServerConfig) (s *Server, err error) {
	s = &Server{
		address: cfg.Address,
	}

	// Инициализация роутинга
	s.router = mux.NewRouter()
	s.router.Handle("/products/stock", s.contentTypeCheck(http.HandlerFunc(s.getStock))).Methods("GET")
	s.router.Handle("/products/reserve", s.contentTypeCheck(http.HandlerFunc(s.reserveProducts))).Methods("POST")
	s.router.Handle("/products/release", s.contentTypeCheck(http.HandlerFunc(s.releaseProducts))).Methods("POST")

	//Инициализация логгера
	logger, err := zap.NewDevelopment()
	if err != nil {
		return
	}
	s.logger = logger.Sugar()

	//Инициализация базы данных
	s.storage, err = database.New(cfg.Database)
	return
}
