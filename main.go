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
	db.CreateTable()

	http.HandleFunc("/play", handlers.PlayHandler)
	http.HandleFunc("/stats", handlers.StatsHandler)
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}
