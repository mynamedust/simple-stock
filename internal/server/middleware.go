package server

import (
	"errors"
	"github.com/google/jsonapi"
	"net/http"
)

func (s *Server) contentTypeCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if contentType := r.Header.Get("Content-Type"); contentType != jsonapi.MediaType {
			err := errors.New("invalid content type expected: " + jsonapi.MediaType)
			s.handleError(w, []error{err}, "content type error", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
