package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Nipun2001M/go-backend-ecommerce/internal/json"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {

	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		log.Println(err)
	}
	json.Write(w, http.StatusOK, products)

}

func (h *handler) GetProductById(w http.ResponseWriter, r *http.Request) {

	idStr :=  chi.URLParam(r, "id") 
    if idStr == "" {
        http.Error(w, "missing product id", http.StatusBadRequest)
        return
    }

    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "invalid product id", http.StatusBadRequest)
        return
    }

	product, err := h.service.GetProductById(r.Context(),int(id))
	if err != nil {
		log.Println(err)
	}
	json.Write(w, http.StatusOK, product)

}
