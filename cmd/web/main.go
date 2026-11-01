package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/form/v4"
	"github.com/kyndverse/catering-management-app/internal/config"
	"github.com/kyndverse/catering-management-app/internal/validator"
)

func main() {
	envConfig := config.NewEnvConfig()
	logger := config.NewLogger(envConfig)
	db := config.NewDB(envConfig)
	mux := http.NewServeMux()
	decoder := form.NewDecoder()
	validator := validator.NewValidator()

	app := &http.Server{
		Addr:    envConfig.Addr,
		Handler: mux,
	}

	config.Bootstrap(&config.BootstrapConfig{
		Logger:    logger,
		Mux:       mux,
		DB:        db,
		Decoder:   decoder,
		Validator: validator,
	})

	fmt.Println("Server running in http://localhost:3000")
	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
