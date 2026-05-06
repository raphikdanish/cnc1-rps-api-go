package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"rps-api-go/game"
	models "rps-api-go/models"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GameRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	req.Choice = strings.ToLower(req.Choice)

	valid := map[string]bool{
		"rock":     true,
		"paper":    true,
		"scissors": true,
	}

	if !valid[req.Choice] {
		http.Error(w, "Invalid choice", http.StatusBadRequest)
		return
	}

	computer := game.GetComputerChoice()
	result := game.DecideWinner(req.Choice, computer)

	resp := models.GameResponse{
		PlayerChoice:   req.Choice,
		ComputerChoice: computer,
		Result:         result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
