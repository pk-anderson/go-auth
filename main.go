package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pk-anderson/go-auth/config"
	"github.com/pk-anderson/go-auth/handlers"
	"github.com/pk-anderson/go-auth/services"
)

func main() {
	cfg := config.LoadConfig("config.yaml")

	authService := services.NewAuthService(cfg.JWT.Secret)
	authHandler := handlers.NewAuthHandler(authService)

	r := mux.NewRouter()
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/validate", authHandler.ValidateToken).Methods("GET")

	fmt.Println("Iniciando servidor na porta 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
