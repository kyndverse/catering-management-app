package router

import (
	"net/http"

	"github.com/kyndverse/catering-management-app/internal/handler"
)

type RouteConfig struct {
	Mux            *http.ServeMux
	PageHandler    *handler.PageHandler
	ProductHandler *handler.ProductHandler
}

func (c *RouteConfig) Setup() {
	fs := http.FileServer(http.Dir("static"))
	c.Mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	c.Mux.HandleFunc("GET /", c.PageHandler.HomePage)
	c.Mux.HandleFunc("GET /products", c.ProductHandler.GetProductsHandler)
	c.Mux.HandleFunc("POST /products", c.ProductHandler.CreateProductHandler)
	c.Mux.HandleFunc("GET /products/create", c.PageHandler.CreateProductForm)
}
