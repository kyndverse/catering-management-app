package config

import (
	"database/sql"

	"golang-web/internal/handler"
	"golang-web/internal/repository"
	"golang-web/internal/router"
	"golang-web/internal/service"
	"golang-web/internal/validator"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	Router    chi.Router
	Logger    *logrus.Logger
	Decoder   *form.Decoder
	Validator *validator.Validator
	DB        *sql.DB
}

func Bootstrap(config *BootstrapConfig) {
	// setup repo
	productRepo := repository.NewProductRepository(config.Logger)

	// setup service
	productService := service.NewProductService(config.Logger, config.DB, config.Validator, productRepo)

	// setup handler
	renderer := handler.NewRenderer()
	pageHandler := handler.NewPageHandler(renderer)
	productRender := handler.NewProductHandler(config.Logger, renderer, config.Decoder, productService)

	// setup route
	routeConfig := &router.RouteConfig{
		Router:         config.Router,
		PageHandler:    pageHandler,
		ProductHandler: productRender,
	}

	routeConfig.Setup()
}
