package main

import (
	"fmt"
	"net/http"

	"github.com/kyndverse/catering-management-app/internal/config"
)

func main() {
	envConfig := config.NewEnvConfig()
	logger := config.NewLogger(envConfig)
	db := config.NewDB(envConfig)
	mux := http.NewServeMux()

	app := &http.Server{
		Addr:    envConfig.Addr,
		Handler: mux,
	}

	config.Bootstrap(&config.BootstrapConfig{
		Logger: logger,
		Mux:    mux,
		DB:     db,
	})

	fmt.Println("Server running in http://localhost:3000")
	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
