package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kyndverse/catering-management-app/internal/handler"
)

type RouteConfig struct {
	Router         chi.Router
	PageHandler    *handler.PageHandler
	ProductHandler *handler.ProductHandler
}

func (c *RouteConfig) Setup() {
	fs := http.FileServer(http.Dir("static"))
	c.Router.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Pages
	c.Router.Get("/", c.PageHandler.HomePage)

	// Products
	c.Router.Route("/products", func(r chi.Router) {
		r.Get("/", c.ProductHandler.GetProductsHandler)
		r.Post("/", c.ProductHandler.CreateProductHandler)
		r.Get("/create", c.PageHandler.CreateProductForm)
	})
}
