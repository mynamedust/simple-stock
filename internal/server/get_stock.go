package server

import (
	"github.com/google/jsonapi"
	"lamodaTest/pkg/models"
	"net/http"
)

type StorehouseCount struct {
	ID    int `jsonapi:"primary,storehouse"`
	Count int `jsonapi:"attr,count"`
}

func (s *Server) GetStock(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != jsonapi.MediaType {
		http.Error(w, "Invalid content type. Expected: application/vnd.api+json", http.StatusBadRequest)
		return
	}

	var data models.Storehouse
	if err := jsonapi.UnmarshalPayload(r.Body, &data); err != nil {
		s.handleError(w, err, "JSONAPI decoding error", http.StatusBadRequest)
		return
	}
	count, err := s.storage.GetStorehouseRemainderByID(data.ID)
	if err != nil {
		s.handleError(w, err, "Database request failed", http.StatusInternalServerError)
		return
	}

	response := StorehouseCount{
		ID:    data.ID,
		Count: count,
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)
	if err := jsonapi.MarshalPayload(w, &response); err != nil {
		s.handleError(w, err, "Response marshalling error", http.StatusInternalServerError)
		return
	}
}
