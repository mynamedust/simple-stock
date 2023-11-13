package server

import (
	"github.com/google/jsonapi"
	"net/http"
	"strconv"
)

func (s *Server) handleError(w http.ResponseWriter, errors []error, message string, statusCode int) {
	var errorsObjects []*jsonapi.ErrorObject

	for _, err := range errors {
		//Логирование ошибки
		s.logger.Errorw(message,
			"error", err.Error(),
		)

		errorsObjects = append(errorsObjects, &jsonapi.ErrorObject{
			Status: strconv.Itoa(statusCode),
			Detail: err.Error(),
		})
	}

	// Сериализуем ошибку в формат JSON:API
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(statusCode)
	jsonapi.MarshalErrors(w, errorsObjects)
}
