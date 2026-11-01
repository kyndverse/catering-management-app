package handler

import (
	"net/http"

	"github.com/kyndverse/catering-management-app/internal/service"
	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	Log            *logrus.Logger
	Renderer       *Renderer
	ProductService *service.ProductService
}

func NewProductHandler(logger *logrus.Logger, renderer *Renderer, productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		Log:            logger,
		Renderer:       renderer,
		ProductService: productService,
	}
}

func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.GetAllProduct(r.Context())
	if err != nil {
		h.Log.WithError(err).Error("error get products")
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"products": products,
	}

	h.Renderer.Render(w, "product-page.html", data)
}
