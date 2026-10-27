package config

import (
	"database/sql"
	"net/http"

	"github.com/kyndverse/catering-management-app/internal/router"
)

type BootstrapConfig struct {
	DB  *sql.DB
	Mux *http.ServeMux
}

func Bootstrap(config *BootstrapConfig) {
	// setup repo

	// setup service

	// setup handler
	
	routeConfig := router.RouteConfig{
		Mux: config.Mux,
		// pass handler
		// pass middleware
	}

	routeConfig.Setup()

	server := &http.Server{
		Addr:    ":3000",
		Handler: config.Mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
