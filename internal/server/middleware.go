package server

import (
	"errors"
	"github.com/google/jsonapi"
	"github.com/mynamedust/simple-stock/pkg/models"
	"io"
	"net/http"
	"reflect"
)

func (s *Server) contentTypeCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if contentType := r.Header.Get("Content-Type"); contentType != jsonapi.MediaType {
			err := errors.New("Invalid content type. Expected: " + jsonapi.MediaType)
			s.handleError(w, err, "Invalid content type", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func parseRequestBody(body io.ReadCloser) (products []models.Product, err error) {
	productsData, err := jsonapi.UnmarshalManyPayload(body, reflect.TypeOf(new(models.Product)))
	if err != nil {
		return
	}
	for _, product := range productsData {
		p, ok := product.(*models.Product)
		if !ok {
			err = errors.New("Invalid request data")
			return
		}
		products = append(products, *p)
	}
	return
}
