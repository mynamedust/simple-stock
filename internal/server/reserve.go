package server

import (
	"github.com/google/jsonapi"
	"lamodaTest/internal/business"
	"lamodaTest/pkg/models"
	"log"
	"net/http"
	"reflect"
)

func (s *Server) ReserveProducts(w http.ResponseWriter, r *http.Request) {
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
	if status, err := s.storage.ReserveProductsByCode(business.GetStocksID(products), productsString, count); err != nil {
		s.handleError(w, err, "Product reservation failed", status)
		return
	}
	w.WriteHeader(http.StatusOK)
}
