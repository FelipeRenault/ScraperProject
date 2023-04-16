package producthandlers

import (
	"PelandoChallenge/pkg/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductHandler struct {
	service *service.Service
}

func NewProductHandler(srvc *service.Service) *ProductHandler {
	return &ProductHandler{
		service: srvc,
	}
}

func (p *ProductHandler) HandleProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	url := r.URL.Query().Get("url")

	product, err := p.service.Product.GetProduct(ctx, url)
	if err != nil {
		fmt.Println("Error fetching product:", err)
		return
	}

	w.Header().Set("Content-type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(product)
}
