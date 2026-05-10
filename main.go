package main

import (
	"net/http"
	"rps-api-go/db"
	"rps-api-go/handlers"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	db.Connect()
	http.HandleFunc("/play", handlers.PlayHandler)
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}
