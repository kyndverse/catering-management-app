package config

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/form/v4"
	"github.com/kyndverse/catering-management-app/internal/handler"
	"github.com/kyndverse/catering-management-app/internal/repository"
	"github.com/kyndverse/catering-management-app/internal/router"
	"github.com/kyndverse/catering-management-app/internal/service"
	"github.com/kyndverse/catering-management-app/internal/validator"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	Mux       *http.ServeMux
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
	routeConfig := router.RouteConfig{
		Mux:            config.Mux,
		PageHandler:    pageHandler,
		ProductHandler: productRender,
	}

	routeConfig.Setup()
}
