package handler

import (
	"net/http"

	"github.com/kyndverse/catering-management-app/internal/model"
)

type PageHandler struct {
	Renderer *Renderer
}

func NewPageHandler(renderer *Renderer) *PageHandler {
	return &PageHandler{
		Renderer: renderer,
	}
}

func (h *PageHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	h.Renderer.Render(w, "home-page.html", nil)
}

func (h *PageHandler) CreateProductForm(w http.ResponseWriter, r *http.Request) {
	productForm := model.FormData[model.CreateProductRequest]{
		Judul:  "Add Products",
		Action: "/products",
		Data:   model.CreateProductRequest{},
		Errors: nil,
	}

	data := map[string]any{
		"title":       "Create Product",
		"ProductForm": productForm,
	}

	h.Renderer.Render(w, "product-form-page.html", data)
}
