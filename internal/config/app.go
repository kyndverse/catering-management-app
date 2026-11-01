package config

import (
	"database/sql"
	"net/http"

	"github.com/kyndverse/catering-management-app/internal/handler"
	"github.com/kyndverse/catering-management-app/internal/repository"
	"github.com/kyndverse/catering-management-app/internal/router"
	"github.com/kyndverse/catering-management-app/internal/service"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	Mux    *http.ServeMux
	Logger *logrus.Logger
	DB     *sql.DB
}

func Bootstrap(config *BootstrapConfig) {
	// setup repo
	productRepo := repository.NewProductRepository(config.Logger)

	// setup service
	productService := service.NewProductService(config.DB, productRepo)

	// setup handler
	renderer := handler.NewRenderer()
	pageHandler := handler.NewPageHandler(renderer)
	productRender := handler.NewProductHandler(config.Logger, renderer, productService)

	// setup route
	routeConfig := router.RouteConfig{
		Mux:            config.Mux,
		PageHandler:    pageHandler,
		ProductHandler: productRender,
	}

	routeConfig.Setup()
}
