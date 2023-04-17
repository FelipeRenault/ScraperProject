package producthandlers

import (
	"errors"
	"net/http"

	libErrors "PelandoChallenge/pkg/lib/errors"
	"PelandoChallenge/pkg/lib/httphelper"
	"PelandoChallenge/pkg/service"
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
		if errors.Is(err, libErrors.ErrSiteNotAvailable) {
			httphelper.BadRequest(w, err.Error())
			return
		}
		httphelper.InternalError(w, err.Error())
		return
	}

	httphelper.RenderJSON(w, product)
}
