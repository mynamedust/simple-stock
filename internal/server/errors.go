package server

import (
	"github.com/google/jsonapi"
	"net/http"
	"strconv"
)

func (s *Server) handleError(w http.ResponseWriter, err error, logMessage string, statusCode int) {
	//Логирование ошибки
	s.logger.Errorw(
		logMessage,
		"error", err.Error(),
	)

	// Сериализуем ошибку в формат JSON:API
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(statusCode)
	jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
		ID:     "1",
		Status: strconv.Itoa(statusCode),
		Title:  logMessage,
		Detail: err.Error(),
	}})
}
