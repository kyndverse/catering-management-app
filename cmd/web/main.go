package main

import (
	"fmt"
	"net/http"

	"golang-web/internal/config"
	"golang-web/internal/validator"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/form/v4"
)

func main() {
	envConfig := config.NewEnvConfig()
	logger := config.NewLogger(envConfig)
	db := config.NewDB(envConfig)
	chiRouter := chi.NewRouter()
	decoder := form.NewDecoder()
	validator := validator.NewValidator()

	app := &http.Server{
		Addr:    envConfig.Addr,
		Handler: chiRouter,
	}

	config.Bootstrap(&config.BootstrapConfig{
		Logger:    logger,
		Router:    chiRouter,
		DB:        db,
		Decoder:   decoder,
		Validator: validator,
	})

	fmt.Println("Server running in http://localhost:3000")
	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
