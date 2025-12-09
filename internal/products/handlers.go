package products

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

func (h *handler) ListProducts(w http.ResponseWriter,r *http.Request) {

	products,err:=h.service.ListProducts(r.Context())
	if err!=nil{
		log.Println(err)
	}
	json.Write(w,http.StatusOK,products)
	


}