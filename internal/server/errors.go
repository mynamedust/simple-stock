package server

import (
	"github.com/google/jsonapi"
	"net/http"
	"strconv"
)

// handleError логирует оишбки из аргумента и выводит их клиенту.
func (s *Server) handleError(w http.ResponseWriter, errors []error, message string, statusCode int) {
	var errorsObjects []*jsonapi.ErrorObject

	for _, err := range errors {
		s.logger.Errorw(message,
			"error", err.Error(),
		)

		errorsObjects = append(errorsObjects, &jsonapi.ErrorObject{
			Status: strconv.Itoa(statusCode),
			Detail: err.Error(),
		})
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(statusCode)
	jsonapi.MarshalErrors(w, errorsObjects)
}
