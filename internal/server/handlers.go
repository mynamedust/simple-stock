package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type StockData struct {
	StorehouseID int `json:"storehouse_id"`
}

type CountResponse struct {
	Count int `json:"count"`
}

type productData struct {
	StockID int    `json:"stock_id"`
	Code    string `json:"code"`
}

func (s *Server) GetStock(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		http.Error(w, "Invalid content type. Expected: application/json", http.StatusBadRequest)
		return
	}

	var data StockData
	//Нужно добавить проверку на валидный json ( проверка на ключи )
	//изменить запрос на quantity - reserved
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		s.logger.Errorw(
			"JSON decode failed",
			"error", err.Error())
		http.Error(w, "JSON decoding error: "+err.Error(), http.StatusBadRequest)
		return
	}
	count, err := s.storage.GetStockRemainderByID(data.StorehouseID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		s.logger.Errorw(
			"Database request failed",
			"error", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := CountResponse{
		Count: count,
	}
	responseData, err := json.Marshal(response)
	if err != nil {
		s.logger.Errorw(
			"Response marshalling error",
			"error", err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func (s *Server) ReserveProducts(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		http.Error(w, "Invalid content type. Expected: application/json", http.StatusBadRequest)
		return
	}

	var products []productData
	// нужно добавить валидацию приходящих данных по ключу и значению
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&products); err != nil {
		s.logger.Errorw(
			"JSON decode failed",
			"error", err.Error())
		http.Error(w, "JSON decoding error: "+err.Error(), http.StatusBadRequest)
		return
	}
	productsString, count := productToString(products)
	if status, err := s.storage.ReserveProductsByCode(getStocks(products), productsString, count); err != nil {
		s.logger.Errorw(
			"Product reservation failed",
			"error", err.Error())
		http.Error(w, "Product reservation failed: "+err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func productToString(products []productData) (string, int) {
	var productsSting string
	unique := make(map[productData]struct{})

	count := 0
	for _, product := range products {
		if _, exist := unique[product]; !exist {
			count++
		}
		unique[product] = struct{}{}
		productsSting += fmt.Sprintf("('%s', %d),", product.Code, product.StockID)
	}
	// Удаление последней запятой
	return productsSting[:len(productsSting)-1], count
}

func getStocks(products []productData) []int {
	uniqueStocks := make(map[int]struct{})
	for _, product := range products {
		uniqueStocks[product.StockID] = struct{}{}
	}

	var stocks []int
	for id := range uniqueStocks {
		stocks = append(stocks, id)
	}

	return stocks
}

func (s *Server) ReleaseProducts(w http.ResponseWriter, r *http.Request) {
	if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
		http.Error(w, "Invalid content type. Expected: application/json", http.StatusBadRequest)
		return
	}

	var products []productData
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&products); err != nil {
		s.logger.Errorw(
			"JSON decode failed",
			"error", err.Error())
		http.Error(w, "JSON decoding error: "+err.Error(), http.StatusBadRequest)
		return
	}
	productsStrign, count := productToString(products)
	if status, err := s.storage.ReleaseProductsByCode(getStocks(products), productsStrign, count); err != nil {
		s.logger.Errorw(
			"Product releasing failed",
			"error", err.Error())
		http.Error(w, "Product releasing failed: "+err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusOK)
}
