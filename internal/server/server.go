package server

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"lamodaTest/internal/config"
	"lamodaTest/pkg/database"
	"log"
	"net/http"
)

type Server struct {
	router  *mux.Router
	storage database.Storage
	address string //`env:"SERVER_ADDRESS"`
	logger  *zap.SugaredLogger
}

func (s *Server) Run() error {
	defer s.logger.Sync()
	defer s.storage.Close()

	s.logger.Infow(
		"Starting server",
		"Address", s.address)
	return http.ListenAndServe(s.address, s.router)
}

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
