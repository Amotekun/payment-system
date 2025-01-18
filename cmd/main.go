package main

import (
	"log"
	"net/http"

	"github.com/Amotekun/payment-system/internal/config"
	"github.com/Amotekun/payment-system/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/api/register", handlers.Register).Methods("POST")

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
