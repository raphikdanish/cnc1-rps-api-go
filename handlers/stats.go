package handlers

import (
	"encoding/json"
	"net/http"

	"rps-api-go/db"
	"rps-api-go/models"
)

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	var stats models.StatsResponse

	// Total games
	err := db.DB.QueryRow(`
		SELECT COUNT(*) FROM games
	`).Scan(&stats.GamesPlayed)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Wins
	err = db.DB.QueryRow(`
		SELECT COUNT(*) FROM games WHERE result = 'win'
	`).Scan(&stats.Wins)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Losses
	err = db.DB.QueryRow(`
		SELECT COUNT(*) FROM games WHERE result = 'lose'
	`).Scan(&stats.Losses)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Draws
	err = db.DB.QueryRow(`
		SELECT COUNT(*) FROM games WHERE result = 'draw'
	`).Scan(&stats.Draws)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Win percentage
	if stats.GamesPlayed > 0 {
		stats.WinPercentage =
			(float64(stats.Wins) / float64(stats.GamesPlayed)) * 100
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}
