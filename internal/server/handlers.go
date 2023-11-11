package server

import (
	"net/http"
)

func (s *Server) ReleaseProducts(w http.ResponseWriter, r *http.Request) {
	//if contentType := r.Header.Get("Content-Type"); contentType != "application/json" {
	//	http.Error(w, "Invalid content type. Expected: application/json", http.StatusBadRequest)
	//	return
	//}
	//
	//var products []ProductData
	//decoder := json.NewDecoder(r.Body)
	//if err := decoder.Decode(&products); err != nil {
	//	s.logger.Errorw(
	//		"JSON decode failed",
	//		"error", err.Error())
	//	http.Error(w, "JSON decoding error: "+err.Error(), http.StatusBadRequest)
	//	return
	//}
	//productsStrign, count := business.ProductToString(products)
	//if status, err := s.storage.ReleaseProductsByCode(business.GetStocksID(products), productsStrign, count); err != nil {
	//	s.logger.Errorw(
	//		"Product releasing failed",
	//		"error", err.Error())
	//	http.Error(w, "Product releasing failed: "+err.Error(), status)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
}
