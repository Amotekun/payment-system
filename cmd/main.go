package main

import (
	"log"
	"net/http"

	"github.com/Amotekun/payment-system/internal/config"
	"github.com/Amotekun/payment-system/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// initialize the database
	config.InitDB()

	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create a new router
	router := mux.NewRouter()

	// public routes
	router.HandleFunc("/api/register", handlers.Register).Methods("POST")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
