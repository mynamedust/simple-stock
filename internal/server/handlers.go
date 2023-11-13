package server

import (
	"github.com/google/jsonapi"
	"github.com/mynamedust/simple-stock/internal/business"
	"github.com/mynamedust/simple-stock/pkg/models"
	"io"
	"net/http"
)

func parseRequestBody(body io.ReadCloser) (models.ProductDtoWithStorehouseID, error) {
	var products models.ProductDtoWithStorehouseID
	if err := jsonapi.UnmarshalPayload(body, &products); err != nil {
		return products, err
	}
	return products, nil
}

// reserveProducts Обработчик HTTP-запросов для резервирования товаров.
func (s *Server) reserveProducts(w http.ResponseWriter, r *http.Request) {
	products, err := parseRequestBody(r.Body)
	if err != nil {
		s.handleError(w, []error{err}, "JSONAPI decoding error", http.StatusBadRequest)
		return
	}

	if err = business.Reserve(products, s.storage, s.logger); err != nil {
		s.handleError(w, []error{err}, "product reservation failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// releaseProducts Обработчик HTTP-запросов для освобождения резерва товаров.
func (s *Server) releaseProducts(w http.ResponseWriter, r *http.Request) {
	products, err := parseRequestBody(r.Body)
	if err != nil {
		s.handleError(w, []error{err}, "JSONAPI decoding error", http.StatusBadRequest)
		return
	}

	if err = business.Release(products, s.storage, s.logger); err != nil {
		s.handleError(w, []error{err}, "product releasing failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// getStock Обработчик HTTP-запросов для получения количества товаров на складе.
func (s *Server) getStock(w http.ResponseWriter, r *http.Request) {
	var storehouse models.Storehouse
	if err := jsonapi.UnmarshalPayload(r.Body, &storehouse); err != nil {
		s.handleError(w, []error{err}, "JSONAPI decoding error", http.StatusBadRequest)
		return
	}

	count, err := business.GetRemainder(storehouse.ID, s.storage, s.logger)
	if err != nil {
		s.handleError(w, []error{err}, "database request failed", http.StatusInternalServerError)
		return
	}
	storehouse.Count = count

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)
	if err := jsonapi.MarshalPayload(w, &storehouse); err != nil {
		s.handleError(w, []error{err}, "response marshalling error", http.StatusInternalServerError)
		return
	}
}
