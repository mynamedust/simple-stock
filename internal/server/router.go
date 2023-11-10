package server

import "github.com/gorilla/mux"

func (s *Server) initRouter() {
	s.router = mux.NewRouter()
	s.router.HandleFunc("/products/stock", s.GetStock).Methods("GET")
	s.router.HandleFunc("/products/reserve", s.ReserveProducts).Methods("POST")
	s.router.HandleFunc("/products/release", s.ReleaseProducts).Methods("POST")
}
