package server

import (
	"github.com/google/jsonapi"
	"log"
	"net/http"
	"reflect"
	"simple-stock/internal/business"
	"simple-stock/pkg/models"
)

// ReleaseProducts Обработчик HTTP-запросов для освобождения резерва товаров.
func (s *Server) ReleaseProducts(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != jsonapi.MediaType {
		http.Error(w, "Invalid content type. Expected: "+jsonapi.MediaType, http.StatusBadRequest)
		return
	}

	productsData, err := jsonapi.UnmarshalManyPayload(r.Body, reflect.TypeOf(new(models.Product)))
	if err != nil {
		s.handleError(w, err, "JSONAPI decoding error", http.StatusBadRequest)
		return
	}
	var products []models.Product
	for _, product := range productsData {
		p, ok := product.(*models.Product)
		if !ok {
			log.Println("not product")
			return
		}
		products = append(products, *p)
	}
	productsString, count := business.ProductToString(products)
	if status, err := s.storage.ReleaseProductsByCode(business.GetStocksID(products), productsString, count); err != nil {
		s.handleError(w, err, "Product releasing failed", status)
		return
	}
	w.WriteHeader(http.StatusOK)
}
