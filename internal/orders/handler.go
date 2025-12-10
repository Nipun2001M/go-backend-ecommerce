package orders

import (
	"log"
	"net/http"

	"github.com/Nipun2001M/go-backend-ecommerce/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	var tempOrder createOrderParams

	if err := json.Read(r, &tempOrder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)

	}
	createdorder, err := h.service.PlaceOrder(r.Context(), tempOrder)
	if err != nil {

		if err == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	json.Write(w, http.StatusCreated, createdorder)

}
