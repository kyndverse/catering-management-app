package handler

import (
	"net/http"

	"golang-web/internal/model"
	"golang-web/internal/service"

	"github.com/go-playground/form/v4"

	"github.com/sirupsen/logrus"
)

type ProductHandler struct {
	Log            *logrus.Logger
	Renderer       *Renderer
	Decoder        *form.Decoder
	ProductService *service.ProductService
}

func NewProductHandler(logger *logrus.Logger, renderer *Renderer, decoder *form.Decoder, productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		Log:            logger,
		Renderer:       renderer,
		Decoder:        decoder,
		ProductService: productService,
	}
}

func (h *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductService.GetAllProduct(r.Context())
	if err != nil {
		h.Log.WithError(err).Error("error get products")
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"title":    "Products",
		"products": products,
	}

	h.Renderer.Render(w, "product-page.html", data)
}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	request := new(model.CreateProductRequest)
	if err := h.Decoder.Decode(request, r.PostForm); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	_, formError, err := h.ProductService.AddProduct(r.Context(), request)
	if err != nil {
		h.Log.WithError(err).Error("error add products")
		http.Error(w, "Failed to add products", http.StatusInternalServerError)
		return
	}

	if formError != nil {
		productForm := model.FormData[model.CreateProductRequest]{
			Judul:  "Create Product",
			Action: "/products",
			Data:   *request,
			Errors: formError,
		}

		data := map[string]any{
			"title":       "Create Product",
			"ProductForm": productForm,
		}

		h.Renderer.Render(w, "product-form-page.html", data)
		return
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
