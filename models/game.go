package models

type GameRequest struct {
	Choice string `json:"choice"`
}

type GameResponse struct {
	PlayerChoice   string `json:"player_choice"`
	ComputerChoice string `json:"computer_choice"`
	Result         string `json:"result"`
}
