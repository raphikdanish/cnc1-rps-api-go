// @title Rock Paper Scissors API
// @version 1.0
// @description Cloud Native Rock Paper Scissors API
// @host localhost:8080
// @BasePath /
package main

import (
	"net/http"

	"rps-api-go/db"
	"rps-api-go/handlers"

	_ "rps-api-go/docs"

	httpSwagger "github.com/swaggo/http-swagger"
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
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", nil)
}
